package shared

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
}

func NewLobbyInfo(id uint32, code string) *LobbyInfo {
	return &LobbyInfo{
		ID:      id,
		Code:    code,
		Players: make([]*ClientInfo, 2),
		State:   LobbyStateEmpty,
	}
}

func (l *LobbyInfo) AddPlayer(player *ClientInfo) {
	if l.IsFull() {
		return
	}

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
