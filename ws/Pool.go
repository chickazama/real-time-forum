package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/repo"

	"github.com/gorilla/websocket"
)

type Pool struct {
	Login     chan *Client
	Logout    chan *Client
	Clients   map[*Client]bool
	Broadcast chan SocketMessage
	Repo      repo.IRepository
}

func NewPool(repo repo.IRepository) *Pool {
	return &Pool{
		Login:     make(chan *Client),
		Logout:    make(chan *Client),
		Clients:   make(map[*Client]bool),
		Broadcast: make(chan SocketMessage),
		Repo:      repo,
	}
}

func (pool *Pool) Run() {
	for {
		select {
		case client := <-pool.Login:
			pool.Clients[client] = true
			joinedClient := User{
				Code: CodeUserLogin,
				Data: UserData{
					ID:       client.ID,
					Nickname: client.Nickname,
				},
			}
			jsonRes, err := json.Marshal(&joinedClient)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			onlineClients := Users{
				Code: CodeListOnlineUsers,
				Data: []UserData{},
			}
			// Notify all clients of a user join
			for client := range pool.Clients {
				nextClient := UserData{
					ID:       client.ID,
					Nickname: client.Nickname,
				}
				onlineClients.Data = append(onlineClients.Data, nextClient)
				client.Conn.WriteJSON(SocketMessage{Type: websocket.TextMessage, Body: string(jsonRes)})
			}
			// Tell this client all connected users
			// msgBody := fmt.Sprintf("Users Online: %d", len(pool.Clients))
			jsonBody, err := json.Marshal(&onlineClients)
			if err != nil {
				log.Fatal(err.Error())
			}
			client.Conn.WriteJSON(SocketMessage{Type: websocket.TextMessage, Body: string(jsonBody)})
		case client := <-pool.Logout:
			disconnectedClient := User{
				Code: CodeUserLogout,
				Data: UserData{
					ID:       client.ID,
					Nickname: client.Nickname,
				},
			}
			jsonRes, err := json.Marshal(&disconnectedClient)
			if err != nil {
				log.Fatal(err.Error())
			}
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// Notify all clients of a user logout
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
			case CodeDirectMessage:
				var dm DirectMessage
				err = json.Unmarshal([]byte(message.Body), &dm)
				if err != nil {
					log.Fatal(err.Error())
				}
				d := dm.Data
				id, err := dal.CreateMessage(d.SenderID, d.TargetID, d.Author, d.Content, d.Timestamp)
				if err != nil {
					log.Fatal(err.Error())
				}
				dm.Data.ID = id
				body, err := json.Marshal(&dm)
				if err != nil {
					log.Fatal(err.Error())
				}
				for client := range pool.Clients {
					if client.ID == dm.Data.SenderID || client.ID == dm.Data.TargetID {
						if err := client.Conn.WriteJSON(SocketMessage{Type: websocket.TextMessage, Body: string(body)}); err != nil {
							fmt.Println(err)
							return
						}
					}
				}
			case CodeNewComment:
				var c Comment
				err = json.Unmarshal([]byte(message.Body), &c)
				if err != nil {
					log.Fatal(err.Error())
				}
				d := c.Data
				id, err := pool.Repo.CreateComment(d.PostID, d.AuthorID, d.Author, d.Content, d.Timestamp)
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
			case CodeNewPost:
				var p Post
				err = json.Unmarshal([]byte(message.Body), &p)
				if err != nil {
					log.Fatal(err.Error())
				}
				d := p.Data
				id, err := dal.CreatePost(d.AuthorID, d.Author, d.Content, d.Categories, d.Timestamp)
				if err != nil {
					log.Fatal(err.Error())
				}
				p.Data.ID = id
				body, err := json.Marshal(&p)
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
