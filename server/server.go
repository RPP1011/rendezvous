package server

import (
	"net"

	shared "github.com/RPP1011/rendezvous/shared"
)

type Server struct {
	shared.DefaultReliablePacketHandler
	// Map of client IDs to client objects
	clients map[uint32]*shared.ClientInfo
	// Map of lobby IDs to lobby objects
	lobbies map[uint32]*shared.LobbyInfo
	// Map of lobby codes to lobby IDs
	lobbyCodes map[string]uint32

	// Tcp listener
	tcpListener *net.TCPListener

	// Udp listener
	udpListener *net.UDPConn
}

func NewServer() *Server {
	tcpAddr, tcp_addr_err := net.ResolveTCPAddr("tcp", ":8080")
	if tcp_addr_err != nil {
		panic(tcp_addr_err)
	}
	tcpListener, tcp_err := net.ListenTCP("tcp", tcpAddr)

	udpAddr, udp_addr_err := net.ResolveUDPAddr("udp", ":8080")
	if udp_addr_err != nil {
		panic(udp_addr_err)
	}
	udpListener, udp_err := net.ListenUDP("udp", udpAddr)

	if tcp_err != nil {
		panic(tcp_err)
	}

	if udp_err != nil {
		panic(udp_err)
	}

	server := &Server{
		clients:     make(map[uint32]*shared.ClientInfo),
		lobbies:     make(map[uint32]*shared.LobbyInfo),
		lobbyCodes:  make(map[string]uint32),
		tcpListener: tcpListener,
		udpListener: udpListener,
	}

	go server.handle_tcp_connections()

	return server
}
