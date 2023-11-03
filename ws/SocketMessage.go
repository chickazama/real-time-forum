package ws

type SocketMessage struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}
