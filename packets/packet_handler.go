package packets

import (
	"google.golang.org/protobuf/proto"
)

type ReliablePacketHandler interface {
	HandlePacket(packet *ReliablePacket)
	handleRegisterPacket(packet *RegisterPacket)
	handleClientInfoPacket(packet *ClientInfoPacket)
	handleCreateLobbyRequestPacket(packet *CreateLobbyRequestPacket)
	handleCreateLobbyResponsePacket(packet *CreateLobbyResponsePacket)
	handleJoinLobbyPacket(packet *JoinLobbyPacket)
	handleLobbyPlayerInfo(packet *LobbyPlayerInfo)
	handleLobbyInfoPacket(packet *LobbyInfoPacket)
	handleUpdateLobbyPacket(packet *UpdateLobbyPacket)
	handleAttemptDirectConnectPacket(packet *AttemptDirectConnectPacket)
	handleReportDirectConnectResultPacket(packet *ReportDirectConnectResultPacket)
	handleAttemptRelayConnectPacket(packet *AttemptRelayConnectPacket)
	handleKickLobbyMemberPacket(packet *KickLobbyMemberPacket)
	handleKickClientPacket(packet *KickClientPacket)
	handleChangeNamePacket(packet *ChangeNamePacket)
	handleChangeLobbyNamePacket(packet *ChangeLobbyNamePacket)
	handleChatMessagePacket(packet *ChatMessagePacket)
	handleLeaveLobbyPacket(packet *LeaveLobbyPacket)
	handleRegisterConnectionPacket(packet *RegisterConnectionPacket)
}

type IsReliablePacket[T isReliablePacket_Value] interface
{
	GetPacket() T 
}

type DefaultReliablePacketHandler struct {}

func (p *DefaultReliablePacketHandler) HandleReliablePacket(packet *ReliablePacket) {
	switch packet.Value.(type) {
	case *ReliablePacket_RegisterPacket:
		p.handleRegisterPacket(packet.GetRegisterPacket())
	case *ReliablePacket_ClientInfoPacket:
		p.handleClientInfoPacket(packet.GetClientInfoPacket())
	case *ReliablePacket_CreateLobbyRequestPacket:
		p.handleCreateLobbyRequestPacket(packet.GetCreateLobbyRequestPacket())
	case *ReliablePacket_CreateLobbyResponsePacket:
		p.handleCreateLobbyResponsePacket(packet.GetCreateLobbyResponsePacket())
	case *ReliablePacket_JoinLobbyPacket:
		p.handleJoinLobbyPacket(packet.GetJoinLobbyPacket())
	case *ReliablePacket_LobbyPlayerInfo:
		p.handleLobbyPlayerInfo(packet.GetLobbyPlayerInfo())
	case *ReliablePacket_LobbyInfoPacket:
		p.handleLobbyInfoPacket(packet.GetLobbyInfoPacket())
	case *ReliablePacket_UpdateLobbyPacket:
		p.handleUpdateLobbyPacket(packet.GetUpdateLobbyPacket())
	case *ReliablePacket_AttemptDirectConnectPacket:
		p.handleAttemptDirectConnectPacket(packet.GetAttemptDirectConnectPacket())
	case *ReliablePacket_ReportDirectConnectResultPacket:
		p.handleReportDirectConnectResultPacket(packet.GetReportDirectConnectResultPacket())
	case *ReliablePacket_AttemptRelayConnectPacket:
		p.handleAttemptRelayConnectPacket(packet.GetAttemptRelayConnectPacket())
	case *ReliablePacket_KickLobbyMemberPacket:
		p.handleKickLobbyMemberPacket(packet.GetKickLobbyMemberPacket())
	case *ReliablePacket_KickClientPacket:
		p.handleKickClientPacket(packet.GetKickClientPacket())
	case *ReliablePacket_ChangeNamePacket:
		p.handleChangeNamePacket(packet.GetChangeNamePacket())
	case *ReliablePacket_ChangeLobbyNamePacket:
		p.handleChangeLobbyNamePacket(packet.GetChangeLobbyNamePacket())
	case *ReliablePacket_ChatMessagePacket:
		p.handleChatMessagePacket(packet.GetChatMessagePacket())
	case *ReliablePacket_LeaveLobbyPacket:
		p.handleLeaveLobbyPacket(packet.GetLeaveLobbyPacket())
	case *ReliablePacket_RegisterConnectionPacket:
		p.handleRegisterConnectionPacket(packet.GetRegisterConnectionPacket())
	}
}

func (p *DefaultReliablePacketHandler) handleRegisterPacket(packet *RegisterPacket) {
	println("Received RegisterPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleClientInfoPacket(packet *ClientInfoPacket) {
	println("Received ClientInfoPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleCreateLobbyRequestPacket(packet *CreateLobbyRequestPacket) {
	println("Received CreateLobbyRequestPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleCreateLobbyResponsePacket(packet *CreateLobbyResponsePacket) {
	println("Received CreateLobbyResponsePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleJoinLobbyPacket(packet *JoinLobbyPacket) {
	println("Received JoinLobbyPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleLobbyPlayerInfo(packet *LobbyPlayerInfo) {
	println("Received LobbyPlayerInfo packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleLobbyInfoPacket(packet *LobbyInfoPacket) {
	println("Received LobbyInfoPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleUpdateLobbyPacket(packet *UpdateLobbyPacket) {
	println("Received UpdateLobbyPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleAttemptDirectConnectPacket(packet *AttemptDirectConnectPacket) {
	println("Received AttemptDirectConnectPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleReportDirectConnectResultPacket(packet *ReportDirectConnectResultPacket) {
	println("Received ReportDirectConnectResultPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleAttemptRelayConnectPacket(packet *AttemptRelayConnectPacket) {
	println("Received AttemptRelayConnectPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleKickLobbyMemberPacket(packet *KickLobbyMemberPacket) {
	println("Received KickLobbyMemberPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleKickClientPacket(packet *KickClientPacket) {
	println("Received KickClientPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleChangeNamePacket(packet *ChangeNamePacket) {
	println("Received ChangeNamePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleChangeLobbyNamePacket(packet *ChangeLobbyNamePacket) {
	println("Received ChangeLobbyNamePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleChatMessagePacket(packet *ChatMessagePacket) {
	println("Received ChatMessagePacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleLeaveLobbyPacket(packet *LeaveLobbyPacket) {
	println("Received LeaveLobbyPacket packet on DefaultReliablePacketHandler")
}

func (p *DefaultReliablePacketHandler) handleRegisterConnectionPacket(packet *RegisterConnectionPacket) {
	println("Received RegisterConnectionPacket packet on DefaultReliablePacketHandler")
}

func (p *RegisterPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_RegisterPacket{ RegisterPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *ClientInfoPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_ClientInfoPacket{ ClientInfoPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *CreateLobbyRequestPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_CreateLobbyRequestPacket{ CreateLobbyRequestPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *CreateLobbyResponsePacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_CreateLobbyResponsePacket{ CreateLobbyResponsePacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *JoinLobbyPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_JoinLobbyPacket{ JoinLobbyPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *LobbyPlayerInfo) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_LobbyPlayerInfo{ LobbyPlayerInfo: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *LobbyInfoPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_LobbyInfoPacket{ LobbyInfoPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *UpdateLobbyPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_UpdateLobbyPacket{ UpdateLobbyPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *AttemptDirectConnectPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_AttemptDirectConnectPacket{ AttemptDirectConnectPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *ReportDirectConnectResultPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_ReportDirectConnectResultPacket{ ReportDirectConnectResultPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *AttemptRelayConnectPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_AttemptRelayConnectPacket{ AttemptRelayConnectPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *KickLobbyMemberPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_KickLobbyMemberPacket{ KickLobbyMemberPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *KickClientPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_KickClientPacket{ KickClientPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *ChangeNamePacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_ChangeNamePacket{ ChangeNamePacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *ChangeLobbyNamePacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_ChangeLobbyNamePacket{ ChangeLobbyNamePacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *ChatMessagePacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_ChatMessagePacket{ ChatMessagePacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *LeaveLobbyPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_LeaveLobbyPacket{ LeaveLobbyPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *RegisterConnectionPacket) GetPacket() []byte {

	packet := &ReliablePacket{
		Value: &ReliablePacket_RegisterConnectionPacket{ RegisterConnectionPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}


type UnreliablePacketHandler interface {
	HandlePacket(packet *UnreliablePacket)
	handleFormConnectionPacket(packet *FormConnectionPacket)
	handlePingPacket(packet *PingPacket)
	handleDataPacket(packet *DataPacket)
}

type DefaultUnreliablePacketHandler struct {}

func (p *DefaultUnreliablePacketHandler) HandleUnreliablePacket(packet *UnreliablePacket) {
	switch packet.Value.(type) {
	case *UnreliablePacket_FormConnectionPacket:
		p.handleFormConnectionPacket(packet.GetFormConnectionPacket())
	case *UnreliablePacket_PingPacket:
		p.handlePingPacket(packet.GetPingPacket())
	case *UnreliablePacket_DataPacket:
		p.handleDataPacket(packet.GetDataPacket())
	}
}

func (p *DefaultUnreliablePacketHandler) handleFormConnectionPacket(packet *FormConnectionPacket) {
	println("Received FormConnectionPacket packet on DefaultUnreliablePacketHandler")
}

func (p *DefaultUnreliablePacketHandler) handlePingPacket(packet *PingPacket) {
	println("Received PingPacket packet on DefaultUnreliablePacketHandler")
}

func (p *DefaultUnreliablePacketHandler) handleDataPacket(packet *DataPacket) {
	println("Received DataPacket packet on DefaultUnreliablePacketHandler")
}


func (p *FormConnectionPacket) GetPacket() []byte {

	packet := &UnreliablePacket{
		Value: &UnreliablePacket_FormConnectionPacket{ FormConnectionPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}


func (p *PingPacket) GetPacket() []byte {

	packet := &UnreliablePacket{
		Value: &UnreliablePacket_PingPacket{ PingPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}


func (p *DataPacket) GetPacket() []byte {

	packet := &UnreliablePacket{
		Value: &UnreliablePacket_DataPacket{ DataPacket: p },
	}
	data, err := proto.Marshal(packet)
	if err != nil {
		panic(err)
	}
	return data
}