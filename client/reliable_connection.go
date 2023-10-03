package client

import (
	"net"

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
	lobby := shared.NewLobbyInfo(uint32(packet.LobbyId), packet.LobbyCode)
	lobby.AddPlayer(&c.clientInfo)
	c.clientInfo.Lobby = lobby
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
	c.attemptUDPConnection(packet.PeerAddress, func(success bool) ([]byte, error) {
		data := packets.ReportDirectConnectResultPacket{
			Success: success,
		}
		return proto.Marshal(&data)
	})
}

func (c *Client) handleAttemptRelayConnectPacket(packet *packets.AttemptRelayConnectPacket) {
	// Attempt to connect to the relay server
	c.attemptUDPConnection(packet.RelayAddress, func(success bool) ([]byte, error) {
		data := packets.ReportRelayConnectResultPacket{
			Success: success,
		}
		return proto.Marshal(&data)
	})
}

type serializer func(success bool) ([]byte, error)

func (c *Client) attemptUDPConnection(remote_addr_str string, serialize serializer) {
	remote_addr, addr_err := net.ResolveUDPAddr("udp", remote_addr_str)
	if addr_err != nil {
		println("Error resolving other client's address")
		return
	} // Create a new UDP connection
	udpConn, err := net.DialUDP("udp", nil, remote_addr)

	if err != nil {
		println("Error connecting to other client directly")
		return
	}

	c.udpConn = udpConn

	// Send a ReportDirectConnectResultPacket to the server
	data, err := serialize(true)

	if err != nil {
		panic(err)
	}

	c.tcpConn.Write(data)
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
