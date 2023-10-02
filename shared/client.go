package shared

import "net"

type ClientInfo struct {
	// Client's ID
	ID uint32
	// Client's name
	Name string
	// Lobby the client is in
	Lobby *LobbyInfo
	// Client's IP address
	UnreliableConnAddr *net.UDPAddr
	// Client's TCP connection
	ReliableConn *net.TCPConn
}

func (c *ClientInfo) UnreliableConnAddrString() string {
	return c.UnreliableConnAddr.String()
}
