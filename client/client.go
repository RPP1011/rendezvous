package client

import (
	"net"

	"github.com/RPP1011/rendezvous/packets"
	"github.com/RPP1011/rendezvous/shared"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	shared.DefaultReliablePacketHandler
	shared.DefaultUnreliablePacketHandler
	clientInfo                    shared.ClientInfo
	events                        ClientEventHandler
	tcpConn                       *net.TCPConn
	udpConn                       *net.UDPConn
	timeDrift                     int64
	serverConnectionCancelChannel chan struct{}
	lobbyConnectionCancelChannel  chan struct{}
}

func NewClient() *Client {
	return &Client{
		clientInfo: shared.ClientInfo{
			ID:    0,
			Name:  "DefaultName",
			Lobby: nil,
		},
		events:                        ClientEventHandler{},
		tcpConn:                       nil,
		udpConn:                       nil,
		serverConnectionCancelChannel: make(chan struct{}),
		lobbyConnectionCancelChannel:  make(chan struct{}),
	}
}

func (c *Client) GetLobbyInfo() *shared.LobbyInfo {
	return c.clientInfo.Lobby
}

func (c *Client) SetName(name string) {
	c.clientInfo.Name = name
	if c.tcpConn != nil && c.clientInfo.Lobby != nil {
		packet := &packets.ChangeNamePacket{
			NewName: name,
		}
		data, err := proto.Marshal(packet)
		if err != nil {
			panic(err)
		}
		c.tcpConn.Write(data)
	}
}

func (c *Client) Connect(addr string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	c.tcpConn, err = net.DialTCP("tcp", nil, tcpAddr)

	go c.handle_tcp_connection(c.serverConnectionCancelChannel)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) JoinLobby(code string) {
	packet := &packets.JoinLobbyPacket{
		LobbyCode: code,
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	c.tcpConn.Write(data)
}

func (c *Client) CreateLobby() {
	packet := &packets.CreateLobbyRequestPacket{
		LobbyName: c.clientInfo.Name + "'s Lobby",
	}

	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	c.tcpConn.Write(data)
}

func (c *Client) LeaveLobby() {
	packet := &packets.LeaveLobbyPacket{}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	c.tcpConn.Write(data)
}

func (c *Client) ChangeLobbyName(newLobbyName string) {
	packet := &packets.ChangeLobbyNamePacket{
		NewName: newLobbyName,
	}

	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	c.tcpConn.Write(data)
}

func (c *Client) Disconnect() {
	c.tcpConn.Close()
}

func (c *Client) SendChatMessage(message string) {
	packet := &packets.ChatMessagePacket{
		SenderName: c.clientInfo.Name,
		Message:    message,
	}

	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	c.tcpConn.Write(data)
}
