package ui

import (
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/ws"
	"net/http"

	"github.com/gorilla/websocket"
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Check correct HTTP method
	if r.Method != http.MethodGet {
		http.Error(w, "bad request.\n", http.StatusBadRequest)
		return
	}
	userID, err := auth.AuthorizeRequest(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	user, err := dal.GetUserByID(userID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	upgrader := websocket.Upgrader{
		ReadBufferSize:  8192,
		WriteBufferSize: 8192,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	client := ws.Client{
		ID:       userID,
		Nickname: user.Nickname,
		Conn:     conn,
		Pool:     ws.ClientPool,
	}
	client.Pool.Login <- &client
	go client.Read()
}
