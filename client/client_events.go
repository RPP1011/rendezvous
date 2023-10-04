package client

type ReceiveChatMessageCallback func(senderName string, message string)
type UpdateLobbyPingCallback func(ping int64)

type ClientEventHandler struct {
	ReceiveChatMessageCallback
	UpdateLobbyPingCallback
}
