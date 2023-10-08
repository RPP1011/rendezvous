package server

import (
	"github.com/RPP1011/rendezvous/shared"
	"math/rand"
	"net"
	"net/netip"
)

type Server struct {
	// Map of client IDs to client objects
	clients map[uint32]*shared.ClientInfo
	// Map of lobby IDs to lobby objects
	lobbies map[uint32]*shared.LobbyInfo
	// Map of network address to client handler
	handlersById            map[uint32]*ClientSessionHandler
	clientsByReliableAddr   map[netip.AddrPort]*ClientSessionHandler
	clientsByUnreliableAddr map[netip.AddrPort]*ClientSessionHandler
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

	udpAddr, udp_addr_err := net.ResolveUDPAddr("udp", ":8081")
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
		clients:                 make(map[uint32]*shared.ClientInfo),
		lobbies:                 make(map[uint32]*shared.LobbyInfo),
		lobbyCodes:              make(map[string]uint32),
		handlersById:            make(map[uint32]*ClientSessionHandler),
		clientsByReliableAddr:   make(map[netip.AddrPort]*ClientSessionHandler),
		clientsByUnreliableAddr: make(map[netip.AddrPort]*ClientSessionHandler),
		tcpListener:             tcpListener,
		udpListener:             udpListener,
	}

	go server.handle_tcp_connections()

	return server
}

// Valid alphabet for lobby codes
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (s *Server) generate_lobby_code() string {
	buffer := make([]byte, 6)
	for i := 0; i < 6; i++ {
		buffer[i] = alphabet[rand.Intn(len(alphabet))]
	}

	// Check if already exists
	if _, ok := s.lobbyCodes[string(buffer)]; ok {

		return s.generate_lobby_code() // Try again, conflicts are unlikely
	}

	return string(buffer)
}
