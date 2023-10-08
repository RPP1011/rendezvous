package server

import (
	"net"
	"net/netip"

	"github.com/RPP1011/rendezvous/packets"
	"github.com/RPP1011/rendezvous/shared"
	"google.golang.org/protobuf/proto"
)

func (s *Server) handle_tcp_connections() {
	for {
		conn, err := s.tcpListener.AcceptTCP()
		if err != nil {
			panic(err)
		}

		sessionHandler := NewClientSessionHandler(s, conn)
		addr_string := conn.RemoteAddr().String()
		addr_port := netip.MustParseAddrPort(addr_string)
		s.clientsByReliableAddr[addr_port] = sessionHandler

		go sessionHandler.handle_tcp_connection()
	}
}

func (s *Server) handle_udp_registrations() {
	for {
		conn, err := s.tcpForUDPListener.AcceptTCP()
		if err != nil {
			panic(err)
		}

		go s.handleRegistration(conn)
	}
}

func (s *Server) handleRegistration(conn *net.TCPConn) {
	defer conn.Close()
	// Create a buffer to read into
	buffer := make([]byte, 1024*4)

	registered_conn := false
	// Read from the connection
	for {
		bytes_read, read_err := conn.Read(buffer)

		if read_err != nil {
			println("Error reading from TCP for UDP connection")
		}

		// Parse the packet
		container_packet := &packets.ReliablePacket{}

		parse_err := proto.Unmarshal(buffer[:bytes_read], container_packet)

		if parse_err != nil {
			println("Error parsing packet")
			continue
		}

		registerPacket := container_packet.GetRegisterConnectionPacket()

		// Handle the packet
		if registerPacket == nil {
			// Print info of sender
			println("Received packet of type", container_packet.String(), "from", conn.RemoteAddr().String())
			continue
		}

		if registered_conn {
			println("Client already registered")
			continue
		}

		// Check if server already has a client for this handler
		if handler, exists := s.handlersById[registerPacket.ClientId]; exists {
			// Register the UDP connection
			udpAddr, err := net.ResolveUDPAddr("udp", conn.RemoteAddr().String())
			if err != nil {
				panic(err)
			}
			udpConn, connErr := net.DialUDP("udp", s.getLocalUDPAddr(), udpAddr)
			if connErr != nil {
				panic(connErr)
			}

			// Add to maps
			addrPort := netip.MustParseAddrPort(udpConn.RemoteAddr().String())
			s.clientsByUnreliableAddr[addrPort] = handler
			handler.UdpConn = udpConn
			handler.Client.ReliableConnToUDP = conn
			registered_conn = true
		}

	}
}

type ClientSessionHandler struct {
	packets.DefaultReliablePacketHandler
	packets.DefaultUnreliablePacketHandler
	Client  *shared.ClientInfo
	Server  *Server
	TcpConn *net.TCPConn
	UdpConn *net.UDPConn
}

func NewClientSessionHandler(s *Server, conn *net.TCPConn) *ClientSessionHandler {
	return &ClientSessionHandler{
		Server:  s,
		TcpConn: conn,
	}
}

func (c *ClientSessionHandler) handle_tcp_connection() {
	defer c.TcpConn.Close()
	// Create a buffer to read into
	buffer := make([]byte, 1024*4)

	// Read from the connection
	for {
		bytes_read, read_err := c.TcpConn.Read(buffer)

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

// Implement packet handlers

func (c *ClientSessionHandler) handleRegisterPacket(packet *packets.RegisterPacket) {
	// Check if server already has a client for this handler
	if c.Client != nil {
		println("Client already registered")
		return
	}

	// Currently registered client count
	clientCount := uint32(len(c.Server.clients))

	// Create a new client
	client := shared.NewClientInfo(clientCount, packet.Username)

	client.ReliableConn = c.TcpConn

	c.Client = client

	// Add to maps
	c.Server.clients[clientCount] = client

	// Send back a register response
	response := &packets.ClientInfoPacket{
		ClientId: clientCount,
	}

	c.TcpConn.Write(response.GetPacket())
}

func (c *ClientSessionHandler) handleJoinLobbyPacket(packet *packets.JoinLobbyPacket) {
	// Check if client is already in a lobby
	if c.Client.Lobby != nil {
		println("Client already in lobby")
		return
	}

	// Check if lobby exists
	lobbyCode := packet.LobbyCode
	lobbyId, lobby_exists := c.Server.lobbyCodes[lobbyCode]
	if !lobby_exists {
		println("Lobby does not exist")
		return
	}

	lobby := c.Server.lobbies[lobbyId]

	// Check if lobby is full
	if lobby.IsFull() {
		println("Lobby is full")
		return
	}

	// Add client to lobby
	lobby.AddPlayer(c.Client)

	// Send all client an update lobby packet
	lobbyUpdatePacketData := (&packets.UpdateLobbyPacket{
		LobbyName:     lobby.Name,
		PlayerOneInfo: lobby.GetPlayerOneInfo(),
		PlayerTwoInfo: lobby.GetPlayerTwoInfo(),
	}).GetPacket()

	for _, client := range lobby.Players {
		if client == nil {
			continue
		}
		client.ReliableConn.Write(lobbyUpdatePacketData)
	}

	// Send back a LobbyInfoPacket response
	response := &packets.LobbyInfoPacket{
		LobbyId:       lobbyId,
		LobbyCode:     lobbyCode,
		LobbyName:     lobby.Name,
		PlayerOneInfo: lobby.GetPlayerOneInfo(),
		PlayerTwoInfo: lobby.GetPlayerTwoInfo(),
	}

	c.TcpConn.Write(response.GetPacket())

	// If lobby is full, request direct connect
	if lobby.IsFull() {
		// Send direct connect request to all players

		for _, client := range lobby.Players {
			clientConn := client.ReliableConn
			peerConn := lobby.GetOtherPlayer(client).ReliableConnToUDP
			if peerConn == nil {
				// Other player has not registered UDP connection yet
				// This is an invalid state, panicking
				panic("Other player has not registered UDP connection yet")
			}

			directConnectRequestPacket := packets.AttemptDirectConnectPacket{
				PeerAddress: peerConn.RemoteAddr().String(),
			}
			clientConn.Write(directConnectRequestPacket.GetPacket())
		}
	}
}

func (c *ClientSessionHandler) handleChatMessagePacket(packet *packets.ChatMessagePacket) {
	// Broadcast the message to the client's lobby
	lobby := c.Client.Lobby
	if lobby == nil {
		return
	}

	// Send chat message to all clients in the lobby
	for _, client := range lobby.Players {
		if client == nil {
			continue
		}

		clientConn := client.ReliableConn
		clientConn.Write(packet.GetPacket())
	}
}

func (c *ClientSessionHandler) handleCreateLobbyRequest(packet *packets.CreateLobbyRequestPacket) {
	if c.Client == nil {
		println("Client not registered")
		return
	}

	// Check if client is already in a lobby
	if c.Client.Lobby != nil {
		println("Client already in lobby")
		return
	}

	//Generate lobby code
	lobbyCode := c.Server.generate_lobby_code()

	// Create a new lobby

	lobbyId := uint32(len(c.Server.lobbies))
	lobby := shared.NewLobbyInfo(lobbyId, packet.LobbyName, lobbyCode)
	c.Server.lobbies[lobbyId] = lobby
	c.Server.lobbyCodes[lobbyCode] = lobbyId

	// Send CreateLobbyResponsePacket to client
	response := &packets.CreateLobbyResponsePacket{
		LobbyCode: lobbyCode,
	}

	c.TcpConn.Write(response.GetPacket())

}

func (c *ClientSessionHandler) handleLeaveLobbyPacket(packet *packets.LeaveLobbyPacket) {
	// Check if client is registered
	if c.Client == nil {
		println("Client not registered")
		return
	}

	// Check if client is in a lobby
	if c.Client.Lobby == nil {
		println("Client not in lobby")
		return
	}

	// Remove client from lobby
	c.Client.Lobby.RemovePlayer(c.Client)
}

func (c *ClientSessionHandler) handleChangeNamePacket(packet *packets.ChangeNamePacket) {
	// Check if client is registered
	if c.Client == nil {
		println("Client not registered")
		return
	}

	// Change client name
	c.Client.Name = packet.NewName
}

func (c *ClientSessionHandler) handleChangeLobbyNamePacket(packet *packets.ChangeLobbyNamePacket) {
	// Check if client is registered
	if c.Client == nil {
		println("Client not registered")
		return
	}

	// Check if client is in a lobby
	if c.Client.Lobby == nil {
		println("Client not in lobby")
		return
	}

	// Change lobby name
	c.Client.Lobby.Name = packet.NewName

	// Send update lobby to all clients in the lobby
	lobby := c.Client.Lobby

	lobbyUpdatePacket := &packets.UpdateLobbyPacket{
		LobbyName:     lobby.Name,
		PlayerOneInfo: lobby.GetPlayerOneInfo(),
		PlayerTwoInfo: lobby.GetPlayerTwoInfo(),
	}

	for _, client := range lobby.Players {
		if client == nil {
			continue
		}

		clientConn := client.ReliableConn
		clientConn.Write(lobbyUpdatePacket.GetPacket())
	}
}

func (c *ClientSessionHandler) handleKickLobbyMemberPacket(packet *packets.KickLobbyMemberPacket) {
	// Check if client is registered
	if c.Client == nil {
		println("Client not registered")
		return
	}

	// Check if client is in a lobby
	if c.Client.Lobby == nil {
		println("Client not in lobby")
		return
	}

	// Check if client is the lobby owner
	if c.Client.Lobby.GetOwner() != c.Client {
		println("Client not lobby owner")
		return
	}

	// Check if client is kicking themselves
	if packet.PlayerId == c.Client.ID {
		println("Client cannot kick themselves")
		return
	}

	// Check if client is kicking a valid player
	clientToKick, exists := c.Server.clients[packet.PlayerId]
	if !exists {
		println("Client to kick does not exist")
		return
	}

	// Check if client to kick is in the lobby
	if clientToKick.Lobby != c.Client.Lobby {
		println("Client to kick is not in the lobby")
		return
	}

	// Remove client from lobby
	c.Client.Lobby.RemovePlayer(clientToKick)

	// Send kick packet to client
	kickPacket := &packets.KickClientPacket{}
	clientToKick.ReliableConn.Write(kickPacket.GetPacket())

	// Send update lobby to all clients in the lobby
	lobby := c.Client.Lobby

	lobbyUpdatePacket := &packets.UpdateLobbyPacket{
		LobbyName:     lobby.Name,
		PlayerOneInfo: lobby.GetPlayerOneInfo(),
		PlayerTwoInfo: lobby.GetPlayerTwoInfo(),
	}

	for _, client := range lobby.Players {
		if client == nil {
			continue
		}

		clientConn := client.ReliableConn
		clientConn.Write(lobbyUpdatePacket.GetPacket())
	}
}

func (c *ClientSessionHandler) handleReportDirectConnectResultPacket(packet *packets.ReportDirectConnectResultPacket) {
	// Check if client is registered
	if c.Client == nil {
		println("Client not registered")
		return
	}

	// Check if client is in a lobby
	if c.Client.Lobby == nil {
		println("Client not in lobby")
		return
	}

	if packet.Success {
		lobby := c.Client.Lobby
		if lobby.State == shared.LobbyStateDirect {
			lobby.State = shared.LobbyStateDuo
		}
		return
	}

	println("Direct connection failed. Falling back to relay connection")
	// Attempt to connect on relay
	lobby := c.Client.Lobby

	// Send relay connect request to all players
	relayConnectRequestPacket := packets.AttemptRelayConnectPacket{
		RelayAddress: c.Server.udpListener.LocalAddr().String(), // This is pointless
	}

	for _, client := range lobby.Players {
		if client == nil {
			continue
		}

		clientConn := client.ReliableConnToUDP
		clientConn.Write(relayConnectRequestPacket.GetPacket())
	}
}
