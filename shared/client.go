package shared

import "net"

type ClientInfo struct {
	// Client's ID
	ID uint32
	// Client's name
	Name string
	// Lobby the client is in
	Lobby *LobbyInfo
	// Last time the client pinged the server(for use in relayed lobbies)
	LastPing int64
	// Client's IP address
	UnreliableConnAddr *net.UDPAddr
	// Client's TCP connection
	ReliableConn *net.TCPConn
}

func NewClientInfo(id uint32, name string) *ClientInfo {
	return &ClientInfo{
		ID:                 0,
		Name:               name,
		Lobby:              nil,
		LastPing:           0,
		UnreliableConnAddr: nil,
		ReliableConn:       nil,
	}
}

func (c *ClientInfo) UnreliableConnAddrString() string {
	return c.UnreliableConnAddr.String()
}
