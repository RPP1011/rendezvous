package client

import (
	"github.com/RPP1011/rendezvous/packets"
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
			c.HandlePacket(container_packet)
		}
	}
}

// Implement packet handlers

// Handle RegisterPacket
func (c *Client) handleRegisterPacket(packet *packets.ReliablePacket_RegisterPacket) {
	println("handleRegisterPacket invoked on Client")
}

// Handle ClientInfoPacket
func (c *Client) handleClientInfoPacket(packet *packets.ReliablePacket_ClientInfoPacket) {
	println("handleClientInfoPacket invoked on Client")
}

// Handle CreateLobbyRequestPacket
func (c *Client) handleCreateLobbyRequestPacket(packet *packets.ReliablePacket_CreateLobbyRequestPacket) {
	println("handleCreateLobbyRequestPacket invoked on Client")
}

// Handle CreateLobbyResponsePacket
func (c *Client) handleCreateLobbyResponsePacket(packet *packets.ReliablePacket_CreateLobbyResponsePacket) {
	println("handleCreateLobbyResponsePacket invoked on Client")
}

// Handle JoinLobbyPacket
func (c *Client) handleJoinLobbyPacket(packet *packets.ReliablePacket_JoinLobbyPacket) {
	println("handleJoinLobbyPacket invoked on Client")
}

// Handle LobbyInfoPacket
func (c *Client) handleLobbyInfoPacket(packet *packets.ReliablePacket_LobbyInfoPacket) {
	println("handleLobbyInfoPacket invoked on Client")
}

// Handle UpdateLobbyPacket
func (c *Client) handleUpdateLobbyPacket(packet *packets.ReliablePacket_UpdateLobbyPacket) {
	println("handleUpdateLobbyPacket invoked on Client")
}

// Handle AttemptDirectConnectPacket
func (c *Client) handleAttemptDirectConnectPacket(packet *packets.ReliablePacket_AttemptDirectConnectPacket) {
	println("handleAttemptDirectConnectPacket invoked on Client")
}

// Handle ReportDirectConnectResultPacket
func (c *Client) handleReportDirectConnectResultPacket(packet *packets.ReliablePacket_ReportDirectConnectResultPacket) {
	println("handleReportDirectConnectResultPacket invoked on Client")
}

// Handle AttemptRelayConnectPacket
func (c *Client) handleAttemptRelayConnectPacket(packet *packets.ReliablePacket_AttemptRelayConnectPacket) {
	println("handleAttemptRelayConnectPacket invoked on Client")
}

// Handle ReportRelayConnectResultPacket
func (c *Client) handleReportRelayConnectResultPacket(packet *packets.ReliablePacket_ReportRelayConnectResultPacket) {
	println("handleReportRelayConnectResultPacket invoked on Client")
}

// Handle KickLobbyMemberPacket
func (c *Client) handleKickLobbyMemberPacket(packet *packets.ReliablePacket_KickLobbyMemberPacket) {
	println("handleKickLobbyMemberPacket invoked on Client")
}

// Handle KickClientPacket
func (c *Client) handleKickClientPacket(packet *packets.ReliablePacket_KickClientPacket) {
	println("handleKickClientPacket invoked on Client")
}
