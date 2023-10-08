package shared

import "github.com/RPP1011/rendezvous/packets"

type LobbyState uint8

const (
	LobbyStateEmpty LobbyState = iota
	LobbyStateSolo
	LobbyStateDuo
	LobbyStateDirect
	LobbyStateRelay
)

const maxLobbyPlayers = 2

type LobbyInfo struct {
	// Lobby's ID
	ID uint32
	// Lobby's Code
	Code string
	// Lobby Players
	Players []*ClientInfo
	// Lobby State
	State LobbyState
	// Lobby Name
	Name string
}

func NewLobbyInfo(id uint32, name string, code string) *LobbyInfo {
	return &LobbyInfo{
		ID:      id,
		Code:    code,
		Players: make([]*ClientInfo, 2),
		State:   LobbyStateEmpty,
		Name:    name,
	}
}

func (l *LobbyInfo) AddPlayer(player *ClientInfo) {
	if l.IsFull() {
		return
	}

	player.Lobby = l
	l.Players = append(l.Players, player)
	l.State = LobbyStateDuo
}

func (l *LobbyInfo) RemovePlayer(player *ClientInfo) {
	// There are only two players in the lobby at most so an if-else is fine
	if l.Players[0] == player {
		l.Players = l.Players[1:]
	} else if l.Players[1] == player {
		l.Players = l.Players[:1]
	} else {
		return // Player not found
	}

	if len(l.Players) == 0 {
		l.State = LobbyStateEmpty
	} else if len(l.Players) == 1 {
		l.State = LobbyStateSolo
	}
}

func (l *LobbyInfo) IsFull() bool {
	return len(l.Players) >= maxLobbyPlayers
}

func (l *LobbyInfo) IsEmpty() bool {
	return len(l.Players) == 0
}

func (l *LobbyInfo) GetOtherPlayer(player *ClientInfo) *ClientInfo {
	if l.Players[0] == player {
		return l.Players[1]
	} else if l.Players[1] == player {
		return l.Players[0]
	} else {
		return nil
	}
}

/*
message LobbyPlayerInfo {
    int32 playerId = 1;
    string playerName = 2;
}*/
func (l *LobbyInfo) GetPlayerOneInfo() *packets.LobbyPlayerInfo {
	if len(l.Players) == 0 {
		return nil
	}

	return &packets.LobbyPlayerInfo{
		PlayerId:   l.Players[0].ID,
		PlayerName: l.Players[0].Name,
	}
}

func (l *LobbyInfo) GetPlayerTwoInfo() *packets.LobbyPlayerInfo {
	if len(l.Players) < 2 {
		return nil
	}

	return &packets.LobbyPlayerInfo{
		PlayerId:   l.Players[1].ID,
		PlayerName: l.Players[1].Name,
	}
}

// Get owner
func (l *LobbyInfo) GetOwner() *ClientInfo {
	if l.IsEmpty() {
		return nil
	}

	return l.Players[0]
}
