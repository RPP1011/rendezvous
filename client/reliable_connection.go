package client

import (
	"net"
	"time"

	"github.com/RPP1011/rendezvous/packets"
	"github.com/RPP1011/rendezvous/shared"
	"google.golang.org/protobuf/proto"
)

func (c *Client) handle_tcp_connection(cancel <-chan struct{}) {
	// Create a buffer to read into
	buffer := make([]byte, 1024*4)

	// Read from the connection
	for {
		select {
		case <-cancel:
			// Cancel signal received, stop the goroutine
			return
		default:
			bytes_read, read_err := c.tcpConn.Read(buffer)

			if read_err != nil {
				println("Error reading from TCP connection")
			}

			// Parse the packet
			container_packet := &packets.ReliablePacket{}

			parse_err := proto.Unmarshal(buffer[:bytes_read], container_packet)

			if parse_err != nil {
				println("Error parsing packet")
				continue
			}

			// Handle the packet
			c.HandleReliablePacket(container_packet)
		}
	}
}

// Implement packet handlers

// Handle ClientInfoPacket
func (c *Client) handleClientInfoPacket(packet *packets.ClientInfoPacket) {
	c.clientInfo.ID = uint32(packet.ClientId)

	go c.formTCPUDPConn(c.tcpConn.RemoteAddr().String(), c.serverRelayAddress)
}

// Handle CreateLobbyResponsePacket
func (c *Client) handleCreateLobbyResponsePacket(packet *packets.CreateLobbyResponsePacket) {
	// Auto join
	join_packet := &packets.JoinLobbyPacket{
		LobbyCode: packet.LobbyCode,
	}

	data, err := proto.Marshal(join_packet)

	if err != nil {
		panic(err)
	}

	c.tcpConn.Write(data)
}

// Handle LobbyInfoPacket
func (c *Client) handleLobbyInfoPacket(packet *packets.LobbyInfoPacket) {
	lobby := shared.NewLobbyInfo(uint32(packet.LobbyId), packet.LobbyName, packet.LobbyCode)
	lobby.AddPlayer(&c.clientInfo)
	c.clientInfo.Lobby = lobby

	go c.update_drift()
	go c.handle_udp_connection()
}

func (c *Client) handleUpdateLobbyPacket(packet *packets.UpdateLobbyPacket) {
	// Update the lobby
	playerOneInfo := packet.PlayerOneInfo
	playerTwoInfo := packet.PlayerTwoInfo

	lobby := c.clientInfo.Lobby
	if playerOneInfo != nil {
		player := lobby.Players[0]
		player.Name = playerOneInfo.PlayerName
		player.ID = uint32(playerOneInfo.PlayerId)
	}

	if playerTwoInfo != nil {
		player := lobby.Players[1]
		player.Name = playerTwoInfo.PlayerName
		player.ID = uint32(playerTwoInfo.PlayerId)
	}
}

func (c *Client) handleAttemptDirectConnectPacket(packet *packets.AttemptDirectConnectPacket) {
	// Attempt to connect to other client
	otherClientAddr := packet.PeerAddress
	// Use TCP to UDP connection to connect to other client
	tcpAddr, err := net.ResolveTCPAddr("tcp", otherClientAddr)
	if err != nil {
		panic(err)
	}

	udpAddr, err := net.ResolveUDPAddr("udp", otherClientAddr)
	if err != nil {
		panic(err)
	}

	localAddr := c.tcpConnForUDP.LocalAddr().(*net.TCPAddr)

	c.tcpConnForUDP, err = net.DialTCP("tcp", localAddr, tcpAddr)
	if err != nil {
		panic(err)
	}

	// set udp connection
	// Convert the TCP connection to a UDP connection
	localUdpAddr, err := net.ResolveUDPAddr("udp", c.tcpConnForUDP.LocalAddr().(*net.TCPAddr).String())
	c.udpConn, err = net.DialUDP("udp", localUdpAddr, udpAddr)

	// Send packet to other client
	data := packets.PingPacket{
		Timestamp: time.Now().UnixNano() + c.timeDrift,
	}
	packet_data, err := proto.Marshal(&data)
	if err != nil {
		panic(err)
	}
	c.udpConn.Write(packet_data)

}

func (c *Client) handleAttemptRelayConnectPacket(packet *packets.AttemptRelayConnectPacket) {
	// Convert the TCP connection to a UDP connection
	localUdpAddr, err := net.ResolveUDPAddr("udp", c.tcpConnForUDP.LocalAddr().(*net.TCPAddr).String())
	remoteUdpAddr, err := net.ResolveUDPAddr("udp", c.serverRelayAddress)
	udpConn, err := net.DialUDP("udp", localUdpAddr, remoteUdpAddr)

	if err != nil {
		println("Error connecting to other client through relay")
		return
	}
	c.udpConn = udpConn

	if err != nil {
		panic(err)
	}

	// Send packet to server, server will relay after interpreting and calculating route latency
	data := packets.PingPacket{
		Timestamp: time.Now().UnixNano() + c.timeDrift,
	}

	c.udpConn.Write(data.GetPacket())

}

func (c *Client) handleKickClientPacket(packet *packets.KickClientPacket) {
	// Kick the client
	c.tcpConn.Close()
	c.tcpConn = nil
	if c.udpConn != nil {
		c.udpConn.Close()
		c.udpConn = nil
	}
}

func (c *Client) handleChatMessagePacket(packet *packets.ChatMessagePacket) {
	if c.events.ReceiveChatMessageCallback != nil {
		c.events.ReceiveChatMessageCallback(packet.SenderName, packet.Message)
	}
}
