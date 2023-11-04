package ws

import (
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/repo"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	Repo repo.IRepository
}

func NewWebSocketHandler(r repo.IRepository) *WebSocketHandler {
	return &WebSocketHandler{Repo: r}
}

func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	user, err := h.Repo.GetUserByID(userID)
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
	client := Client{
		ID:       userID,
		Nickname: user.Nickname,
		Conn:     conn,
		Pool:     ClientPool,
	}
	client.Pool.Login <- &client
	go client.Read()
}
