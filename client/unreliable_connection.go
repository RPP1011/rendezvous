package client

import (
	"time"

	"github.com/RPP1011/rendezvous/packets"
	"google.golang.org/protobuf/proto"
)

func (c *Client) handle_udp_connection() {
	// Create a buffer to read into
	buffer := make([]byte, 1024*4)

	// Read from the connection
	select {
	case <-c.serverConnectionCancelChannel:
		// Cancel signal received, stop the goroutine
		return
	case <-c.lobbyConnectionCancelChannel:
		// Cancel signal received, stop the goroutine
		return
	default:
		for {
			bytes_read, read_err := c.udpConn.Read(buffer)

			if read_err != nil {
				println("Error reading from UDP connection")
			}

			// Parse the packet
			container_packet := &packets.UnreliablePacket{}

			parse_err := proto.Unmarshal(buffer[:bytes_read], container_packet)

			if parse_err != nil {
				println("Error parsing packet")
				continue
			}

			// Handle the packet
			go c.HandleUnreliablePacket(container_packet)
		}
	}
}

func (c *Client) handlePingPacket(packet *packets.PingPacket) {
	// Convert to adjusted time using time drift

	currentTime := time.Now().UnixNano() + c.timeDrift
	difference := currentTime - packet.Timestamp

	// Calculate the round trip time
	roundTripTime := difference * 2

	// Call event handler
	if c.events.UpdateLobbyPingCallback != nil {
		c.events.UpdateLobbyPingCallback(roundTripTime)
	}
}
