package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/transport"

	"github.com/gorilla/websocket"
)

type Pool struct {
	Login     chan *Client
	Logout    chan *Client
	Clients   map[*Client]bool
	Broadcast chan SocketMessage
}

func NewPool() *Pool {
	return &Pool{
		Login:     make(chan *Client),
		Logout:    make(chan *Client),
		Clients:   make(map[*Client]bool),
		Broadcast: make(chan SocketMessage),
	}
}

func (pool *Pool) Run() {
	for {
		select {
		case client := <-pool.Login:
			pool.Clients[client] = true
			joinedClient := transport.UserResponse{
				ID:       client.ID,
				Nickname: client.Nickname,
			}
			jsonRes, err := json.Marshal(&joinedClient)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			var onlineList []transport.UserResponse
			// Notify all clients of a user join
			for client := range pool.Clients {
				nextClient := transport.UserResponse{
					ID:       client.ID,
					Nickname: client.Nickname,
				}
				onlineList = append(onlineList, nextClient)
				client.Conn.WriteJSON(SocketMessage{Type: websocket.TextMessage, Body: string(jsonRes)})
			}
			// Tell this client all connected users
			// msgBody := fmt.Sprintf("Users Online: %d", len(pool.Clients))
			jsonBody, err := json.Marshal(&onlineList)
			if err != nil {
				log.Fatal(err.Error())
			}
			client.Conn.WriteJSON(SocketMessage{Type: websocket.TextMessage, Body: string(jsonBody)})
		case client := <-pool.Logout:
			loggedOutClient := transport.UserResponse{
				ID:       client.ID,
				Nickname: client.Nickname,
			}
			jsonRes, err := json.Marshal(&loggedOutClient)
			if err != nil {
				log.Fatal(err.Error())
			}
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(SocketMessage{Type: websocket.TextMessage, Body: string(jsonRes)})
			}
		case message := <-pool.Broadcast:
			var body MessageBody
			err := json.Unmarshal([]byte(message.Body), &body)
			if err != nil {
				log.Fatal(err.Error())
			}
			switch body.Code {
			case CodeNewComment:
				var c Comment
				err = json.Unmarshal([]byte(message.Body), &c)
				if err != nil {
					log.Fatal(err.Error())
				}
				d := c.Data
				id, err := dal.CreateComment(d.PostID, d.AuthorID, d.Author, d.Content, d.Timestamp)
				if err != nil {
					log.Fatal(err.Error())
				}
				c.Data.ID = id
				body, err := json.Marshal(&c)
				if err != nil {
					log.Fatal(err.Error())
				}
				for client := range pool.Clients {
					if err := client.Conn.WriteJSON(SocketMessage{Type: websocket.TextMessage, Body: string(body)}); err != nil {
						fmt.Println(err)
						return
					}
				}

			}
		}
	}
}
