package shared

import (
	"github.com/RPP1011/rendezvous/packets"
)

type ReliablePacketHandler interface {
	HandlePacket(packet *packets.ReliablePacket)
	handleRegisterPacket(packet *packets.ReliablePacket_RegisterPacket)
	handleClientInfoPacket(packet *packets.ReliablePacket_ClientInfoPacket)
	handleCreateLobbyRequestPacket(packet *packets.ReliablePacket_CreateLobbyRequestPacket)
	handleCreateLobbyResponsePacket(packet *packets.ReliablePacket_CreateLobbyResponsePacket)
	handleJoinLobbyPacket(packet *packets.ReliablePacket_JoinLobbyPacket)
	handleLobbyInfoPacket(packet *packets.ReliablePacket_LobbyInfoPacket)
	handleUpdateLobbyPacket(packet *packets.ReliablePacket_UpdateLobbyPacket)
	handleAttemptDirectConnectPacket(packet *packets.ReliablePacket_AttemptDirectConnectPacket)
	handleReportDirectConnectResultPacket(packet *packets.ReliablePacket_ReportDirectConnectResultPacket)
	handleAttemptRelayConnectPacket(packet *packets.ReliablePacket_AttemptRelayConnectPacket)
	handleReportRelayConnectResultPacket(packet *packets.ReliablePacket_ReportRelayConnectResultPacket)
	handleKickLobbyMemberPacket(packet *packets.ReliablePacket_KickLobbyMemberPacket)
	handleKickClientPacket(packet *packets.ReliablePacket_KickClientPacket)
	handleChangeNamePacket(packet *packets.ReliablePacket_ChangeNamePacket)
	handleChangeLobbyNamePacket(packet *packets.ReliablePacket_ChangeLobbyNamePacket)
	handleChatMessagePacket(packet *packets.ReliablePacket_ChatMessagePacket)
}

type DefaultReliablePacketHandler struct {}

func (p *DefaultReliablePacketHandler) HandlePacket(packet *packets.ReliablePacket) {
	switch packet.Value.(type) {
	case *packets.ReliablePacket_RegisterPacket:
		p.handleRegisterPacket(packet.Value.(*packets.ReliablePacket_RegisterPacket))
	case *packets.ReliablePacket_ClientInfoPacket:
		p.handleClientInfoPacket(packet.Value.(*packets.ReliablePacket_ClientInfoPacket))
	case *packets.ReliablePacket_CreateLobbyRequestPacket:
		p.handleCreateLobbyRequestPacket(packet.Value.(*packets.ReliablePacket_CreateLobbyRequestPacket))
	case *packets.ReliablePacket_CreateLobbyResponsePacket:
		p.handleCreateLobbyResponsePacket(packet.Value.(*packets.ReliablePacket_CreateLobbyResponsePacket))
	case *packets.ReliablePacket_JoinLobbyPacket:
		p.handleJoinLobbyPacket(packet.Value.(*packets.ReliablePacket_JoinLobbyPacket))
	case *packets.ReliablePacket_LobbyInfoPacket:
		p.handleLobbyInfoPacket(packet.Value.(*packets.ReliablePacket_LobbyInfoPacket))
	case *packets.ReliablePacket_UpdateLobbyPacket:
		p.handleUpdateLobbyPacket(packet.Value.(*packets.ReliablePacket_UpdateLobbyPacket))
	case *packets.ReliablePacket_AttemptDirectConnectPacket:
		p.handleAttemptDirectConnectPacket(packet.Value.(*packets.ReliablePacket_AttemptDirectConnectPacket))
	case *packets.ReliablePacket_ReportDirectConnectResultPacket:
		p.handleReportDirectConnectResultPacket(packet.Value.(*packets.ReliablePacket_ReportDirectConnectResultPacket))
	case *packets.ReliablePacket_AttemptRelayConnectPacket:
		p.handleAttemptRelayConnectPacket(packet.Value.(*packets.ReliablePacket_AttemptRelayConnectPacket))
	case *packets.ReliablePacket_ReportRelayConnectResultPacket:
		p.handleReportRelayConnectResultPacket(packet.Value.(*packets.ReliablePacket_ReportRelayConnectResultPacket))
	case *packets.ReliablePacket_KickLobbyMemberPacket:
		p.handleKickLobbyMemberPacket(packet.Value.(*packets.ReliablePacket_KickLobbyMemberPacket))
	case *packets.ReliablePacket_KickClientPacket:
		p.handleKickClientPacket(packet.Value.(*packets.ReliablePacket_KickClientPacket))
	case *packets.ReliablePacket_ChangeNamePacket:
		p.handleChangeNamePacket(packet.Value.(*packets.ReliablePacket_ChangeNamePacket))
	case *packets.ReliablePacket_ChangeLobbyNamePacket:
		p.handleChangeLobbyNamePacket(packet.Value.(*packets.ReliablePacket_ChangeLobbyNamePacket))
	case *packets.ReliablePacket_ChatMessagePacket:
		p.handleChatMessagePacket(packet.Value.(*packets.ReliablePacket_ChatMessagePacket))
	}
}

func (p *DefaultReliablePacketHandler) handleRegisterPacket(packet *packets.ReliablePacket_RegisterPacket) {
	println("Received RegisterPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleClientInfoPacket(packet *packets.ReliablePacket_ClientInfoPacket) {
	println("Received ClientInfoPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleCreateLobbyRequestPacket(packet *packets.ReliablePacket_CreateLobbyRequestPacket) {
	println("Received CreateLobbyRequestPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleCreateLobbyResponsePacket(packet *packets.ReliablePacket_CreateLobbyResponsePacket) {
	println("Received CreateLobbyResponsePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleJoinLobbyPacket(packet *packets.ReliablePacket_JoinLobbyPacket) {
	println("Received JoinLobbyPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleLobbyInfoPacket(packet *packets.ReliablePacket_LobbyInfoPacket) {
	println("Received LobbyInfoPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleUpdateLobbyPacket(packet *packets.ReliablePacket_UpdateLobbyPacket) {
	println("Received UpdateLobbyPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleAttemptDirectConnectPacket(packet *packets.ReliablePacket_AttemptDirectConnectPacket) {
	println("Received AttemptDirectConnectPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleReportDirectConnectResultPacket(packet *packets.ReliablePacket_ReportDirectConnectResultPacket) {
	println("Received ReportDirectConnectResultPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleAttemptRelayConnectPacket(packet *packets.ReliablePacket_AttemptRelayConnectPacket) {
	println("Received AttemptRelayConnectPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleReportRelayConnectResultPacket(packet *packets.ReliablePacket_ReportRelayConnectResultPacket) {
	println("Received ReportRelayConnectResultPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleKickLobbyMemberPacket(packet *packets.ReliablePacket_KickLobbyMemberPacket) {
	println("Received KickLobbyMemberPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleKickClientPacket(packet *packets.ReliablePacket_KickClientPacket) {
	println("Received KickClientPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleChangeNamePacket(packet *packets.ReliablePacket_ChangeNamePacket) {
	println("Received ChangeNamePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleChangeLobbyNamePacket(packet *packets.ReliablePacket_ChangeLobbyNamePacket) {
	println("Received ChangeLobbyNamePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleChatMessagePacket(packet *packets.ReliablePacket_ChatMessagePacket) {
	println("Received ChatMessagePacket packet on DefaultReliablePacketHandler")
}
