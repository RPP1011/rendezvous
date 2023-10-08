package server

import (
	"github.com/RPP1011/rendezvous/packets"
	"google.golang.org/protobuf/proto"
	"net"
)

func (s *Server) handle_udp_messages() {
	buffer := make([]byte, 1024*4)
	for {
		n, addr, err := s.udpListener.ReadFromUDPAddrPort(buffer)
		if err != nil {
			println("Error reading from UDP connection")
		}

		// Parse the packet
		container_packet := &packets.UnreliablePacket{}
		proto.Unmarshal(buffer[:n], container_packet)

		// get client handler if it is in server already or create a new one
		handler, ok := s.clientsByUnreliableAddr[addr]

		if !ok {
			formConnectionPacket := container_packet.GetFormConnectionPacket()
			// Check if the client is a form connection request, if not, ignore
			if formConnectionPacket == nil {
				continue
			}

			// Get current client handler for client
			handler = s.handlersById[formConnectionPacket.ClientId]
			handler.Client.UnreliableConnAddr, err = net.ResolveUDPAddr("udp", addr.String())
			if err != nil {
				println("Error resolving UDP address")
			}

			// Add conn to mapping
			s.clientsByUnreliableAddr[addr] = handler

		}

		// Handle the packet
		go handler.HandleUnreliablePacket(container_packet)
	}
}

// Implement packet handlers
func (c *ClientSessionHandler) handlePingPacket(packet *packets.PingPacket) {
	client := c.Client
	client.LastPing = packet.Timestamp

	lobby := client.Lobby
	peer := lobby.GetOtherPlayer(client)

	// Sum lobby pings
	sum := packet.Timestamp + peer.LastPing

	// Send the ping to peer

	ping_packet := &packets.PingPacket{
		Timestamp: sum,
	}

	c.Server.udpListener.WriteToUDPAddrPort(ping_packet.GetPacket(), peer.UnreliableConnAddr.AddrPort())
}
