package client

type ReceiveChatMessageCallback func(senderName string, message string)

type ClientEventHandler struct {
	ReceiveChatMessageCallback
}
