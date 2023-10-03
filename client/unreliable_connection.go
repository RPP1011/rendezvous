package client

import (
	"github.com/RPP1011/rendezvous/packets"
	"google.golang.org/protobuf/proto"
)

func (c *Client) handle_udp_connection() {
	// Create a buffer to read into
	buffer := make([]byte, 1024*4)

	// Read from the connection
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
		c.HandleUnreliablePacket(container_packet)
	}
}
