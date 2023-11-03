package models

type Message struct {
	ID        int    `json:"id"`
	SenderID  int    `json:"senderID"`
	TargetID  int    `json:"targetID"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}
