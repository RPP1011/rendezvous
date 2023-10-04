package shared

import (
	"github.com/RPP1011/rendezvous/packets"
)

type ReliablePacketHandler interface {
	HandlePacket(packet *packets.ReliablePacket)
	handleRegisterPacket(packet *packets.RegisterPacket)
	handleClientInfoPacket(packet *packets.ClientInfoPacket)
	handleCreateLobbyRequestPacket(packet *packets.CreateLobbyRequestPacket)
	handleCreateLobbyResponsePacket(packet *packets.CreateLobbyResponsePacket)
	handleJoinLobbyPacket(packet *packets.JoinLobbyPacket)
	handleLobbyPlayerInfo(packet *packets.LobbyPlayerInfo)
	handleLobbyInfoPacket(packet *packets.LobbyInfoPacket)
	handleUpdateLobbyPacket(packet *packets.UpdateLobbyPacket)
	handleAttemptDirectConnectPacket(packet *packets.AttemptDirectConnectPacket)
	handleReportDirectConnectResultPacket(packet *packets.ReportDirectConnectResultPacket)
	handleAttemptRelayConnectPacket(packet *packets.AttemptRelayConnectPacket)
	handleReportRelayConnectResultPacket(packet *packets.ReportRelayConnectResultPacket)
	handleKickLobbyMemberPacket(packet *packets.KickLobbyMemberPacket)
	handleKickClientPacket(packet *packets.KickClientPacket)
	handleChangeNamePacket(packet *packets.ChangeNamePacket)
	handleChangeLobbyNamePacket(packet *packets.ChangeLobbyNamePacket)
	handleChatMessagePacket(packet *packets.ChatMessagePacket)
	handleLeaveLobbyPacket(packet *packets.LeaveLobbyPacket)
}

type DefaultReliablePacketHandler struct {}

func (p *DefaultReliablePacketHandler) HandleReliablePacket(packet *packets.ReliablePacket) {
	switch packet.Value.(type) {
	case *packets.ReliablePacket_RegisterPacket:
		p.handleRegisterPacket(packet.GetRegisterPacket())
	case *packets.ReliablePacket_ClientInfoPacket:
		p.handleClientInfoPacket(packet.GetClientInfoPacket())
	case *packets.ReliablePacket_CreateLobbyRequestPacket:
		p.handleCreateLobbyRequestPacket(packet.GetCreateLobbyRequestPacket())
	case *packets.ReliablePacket_CreateLobbyResponsePacket:
		p.handleCreateLobbyResponsePacket(packet.GetCreateLobbyResponsePacket())
	case *packets.ReliablePacket_JoinLobbyPacket:
		p.handleJoinLobbyPacket(packet.GetJoinLobbyPacket())
	case *packets.ReliablePacket_LobbyPlayerInfo:
		p.handleLobbyPlayerInfo(packet.GetLobbyPlayerInfo())
	case *packets.ReliablePacket_LobbyInfoPacket:
		p.handleLobbyInfoPacket(packet.GetLobbyInfoPacket())
	case *packets.ReliablePacket_UpdateLobbyPacket:
		p.handleUpdateLobbyPacket(packet.GetUpdateLobbyPacket())
	case *packets.ReliablePacket_AttemptDirectConnectPacket:
		p.handleAttemptDirectConnectPacket(packet.GetAttemptDirectConnectPacket())
	case *packets.ReliablePacket_ReportDirectConnectResultPacket:
		p.handleReportDirectConnectResultPacket(packet.GetReportDirectConnectResultPacket())
	case *packets.ReliablePacket_AttemptRelayConnectPacket:
		p.handleAttemptRelayConnectPacket(packet.GetAttemptRelayConnectPacket())
	case *packets.ReliablePacket_ReportRelayConnectResultPacket:
		p.handleReportRelayConnectResultPacket(packet.GetReportRelayConnectResultPacket())
	case *packets.ReliablePacket_KickLobbyMemberPacket:
		p.handleKickLobbyMemberPacket(packet.GetKickLobbyMemberPacket())
	case *packets.ReliablePacket_KickClientPacket:
		p.handleKickClientPacket(packet.GetKickClientPacket())
	case *packets.ReliablePacket_ChangeNamePacket:
		p.handleChangeNamePacket(packet.GetChangeNamePacket())
	case *packets.ReliablePacket_ChangeLobbyNamePacket:
		p.handleChangeLobbyNamePacket(packet.GetChangeLobbyNamePacket())
	case *packets.ReliablePacket_ChatMessagePacket:
		p.handleChatMessagePacket(packet.GetChatMessagePacket())
	case *packets.ReliablePacket_LeaveLobbyPacket:
		p.handleLeaveLobbyPacket(packet.GetLeaveLobbyPacket())
	}
}

func (p *DefaultReliablePacketHandler) handleRegisterPacket(packet *packets.RegisterPacket) {
	println("Received RegisterPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleClientInfoPacket(packet *packets.ClientInfoPacket) {
	println("Received ClientInfoPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleCreateLobbyRequestPacket(packet *packets.CreateLobbyRequestPacket) {
	println("Received CreateLobbyRequestPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleCreateLobbyResponsePacket(packet *packets.CreateLobbyResponsePacket) {
	println("Received CreateLobbyResponsePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleJoinLobbyPacket(packet *packets.JoinLobbyPacket) {
	println("Received JoinLobbyPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleLobbyPlayerInfo(packet *packets.LobbyPlayerInfo) {
	println("Received LobbyPlayerInfo packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleLobbyInfoPacket(packet *packets.LobbyInfoPacket) {
	println("Received LobbyInfoPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleUpdateLobbyPacket(packet *packets.UpdateLobbyPacket) {
	println("Received UpdateLobbyPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleAttemptDirectConnectPacket(packet *packets.AttemptDirectConnectPacket) {
	println("Received AttemptDirectConnectPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleReportDirectConnectResultPacket(packet *packets.ReportDirectConnectResultPacket) {
	println("Received ReportDirectConnectResultPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleAttemptRelayConnectPacket(packet *packets.AttemptRelayConnectPacket) {
	println("Received AttemptRelayConnectPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleReportRelayConnectResultPacket(packet *packets.ReportRelayConnectResultPacket) {
	println("Received ReportRelayConnectResultPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleKickLobbyMemberPacket(packet *packets.KickLobbyMemberPacket) {
	println("Received KickLobbyMemberPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleKickClientPacket(packet *packets.KickClientPacket) {
	println("Received KickClientPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleChangeNamePacket(packet *packets.ChangeNamePacket) {
	println("Received ChangeNamePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleChangeLobbyNamePacket(packet *packets.ChangeLobbyNamePacket) {
	println("Received ChangeLobbyNamePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleChatMessagePacket(packet *packets.ChatMessagePacket) {
	println("Received ChatMessagePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleLeaveLobbyPacket(packet *packets.LeaveLobbyPacket) {
	println("Received LeaveLobbyPacket packet on DefaultReliablePacketHandler")
}



type UnreliablePacketHandler interface {
	HandlePacket(packet *packets.UnreliablePacket)
	handlePingPacket(packet *packets.PingPacket)
	handleDataPacket(packet *packets.DataPacket)
}

type DefaultUnreliablePacketHandler struct {}

func (p *DefaultUnreliablePacketHandler) HandleUnreliablePacket(packet *packets.UnreliablePacket) {
	switch packet.Value.(type) {
	case *packets.UnreliablePacket_PingPacket:
		p.handlePingPacket(packet.GetPingPacket())
	case *packets.UnreliablePacket_DataPacket:
		p.handleDataPacket(packet.GetDataPacket())
	}
}

func (p *DefaultUnreliablePacketHandler) handlePingPacket(packet *packets.PingPacket) {
	println("Received PingPacket packet on DefaultUnreliablePacketHandler")
}

func (p *DefaultUnreliablePacketHandler) handleDataPacket(packet *packets.DataPacket) {
	println("Received DataPacket packet on DefaultUnreliablePacketHandler")
}
