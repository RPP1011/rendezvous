package server

import (
	"net"

	"github.com/RPP1011/rendezvous/packets"
	"google.golang.org/protobuf/proto"
)

func (s *Server) handle_tcp_connections() {
	for {
		conn, err := s.tcpListener.AcceptTCP()
		if err != nil {
			panic(err)
		}

		go s.handle_tcp_connection(conn)
	}
}

func (s *Server) handle_tcp_connection(conn *net.TCPConn) {
	defer conn.Close()
	// Create a buffer to read into
	buffer := make([]byte, 1024*4)

	// Read from the connection
	for {
		bytes_read, read_err := conn.Read(buffer)

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
		s.HandleReliablePacket(container_packet)
	}
}

// Implement packet handlers

// Handle RegisterPacket
func (s *Server) handleRegisterPacket(packet *packets.ReliablePacket_RegisterPacket) {
	println("handleRegisterPacket invoked on Server")
}

// Handle ClientInfoPacket
func (s *Server) handleClientInfoPacket(packet *packets.ReliablePacket_ClientInfoPacket) {
	println("handleClientInfoPacket invoked on Server")
}

// Handle CreateLobbyRequestPacket
func (s *Server) handleCreateLobbyRequestPacket(packet *packets.ReliablePacket_CreateLobbyRequestPacket) {
	println("handleCreateLobbyRequestPacket invoked on Server")
}

// Handle CreateLobbyResponsePacket
func (s *Server) handleCreateLobbyResponsePacket(packet *packets.ReliablePacket_CreateLobbyResponsePacket) {
	println("handleCreateLobbyResponsePacket invoked on Server")
}

// Handle JoinLobbyPacket
func (s *Server) handleJoinLobbyPacket(packet *packets.ReliablePacket_JoinLobbyPacket) {
	println("handleJoinLobbyPacket invoked on Server")
}

// Handle LobbyInfoPacket
func (s *Server) handleLobbyInfoPacket(packet *packets.ReliablePacket_LobbyInfoPacket) {
	println("handleLobbyInfoPacket invoked on Server")
}

// Handle UpdateLobbyPacket
func (s *Server) handleUpdateLobbyPacket(packet *packets.ReliablePacket_UpdateLobbyPacket) {
	println("handleUpdateLobbyPacket invoked on Server")
}

// Handle AttemptDirectConnectPacket
func (s *Server) handleAttemptDirectConnectPacket(packet *packets.ReliablePacket_AttemptDirectConnectPacket) {
	println("handleAttemptDirectConnectPacket invoked on Server")
}

// Handle ReportDirectConnectResultPacket
func (s *Server) handleReportDirectConnectResultPacket(packet *packets.ReliablePacket_ReportDirectConnectResultPacket) {
	println("handleReportDirectConnectResultPacket invoked on Server")
}

// Handle AttemptRelayConnectPacket
func (s *Server) handleAttemptRelayConnectPacket(packet *packets.ReliablePacket_AttemptRelayConnectPacket) {
	println("handleAttemptRelayConnectPacket invoked on Server")
}

// Handle ReportRelayConnectResultPacket
func (s *Server) handleReportRelayConnectResultPacket(packet *packets.ReliablePacket_ReportRelayConnectResultPacket) {
	println("handleReportRelayConnectResultPacket invoked on Server")
}

// Handle KickLobbyMemberPacket
func (s *Server) handleKickLobbyMemberPacket(packet *packets.ReliablePacket_KickLobbyMemberPacket) {
	println("handleKickLobbyMemberPacket invoked on Server")
}

// Handle KickClientPacket
func (s *Server) handleKickClientPacket(packet *packets.ReliablePacket_KickClientPacket) {
	println("handleKickClientPacket invoked on Server")
}
