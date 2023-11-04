package api

import (
	"encoding/json"
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/repo"
	"net/http"
	"strconv"
)

type ChatHandler struct {
	Repo repo.IRepository
}

func NewChatHandler(r repo.IRepository) *ChatHandler {
	return &ChatHandler{Repo: r}
}

func (h *ChatHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.AuthorizeRequest(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	err = r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	targetID, err := strconv.Atoi(r.PostFormValue("targetID"))
	if err != nil {
		log.Println("here")
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	messages, err := h.Repo.GetMessagesBySenderAndTargetIDs(userID, targetID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&messages)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
}
