package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID       int
	Nickname string
	Conn     *websocket.Conn
	Pool     *Pool
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Logout <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := SocketMessage{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		// c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
