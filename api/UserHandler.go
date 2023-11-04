package api

import (
	"encoding/json"
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/repo"
	"matthewhope/real-time-forum/transport"
	"net/http"
)

type UserHandler struct {
	Repo repo.IRepository
}

func NewUserHandler(r repo.IRepository) *UserHandler {
	return &UserHandler{Repo: r}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.AuthorizeRequest(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	user, err := h.Repo.GetUserByID(userID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error retrieving users from database.\n", http.StatusInternalServerError)
		return
	}
	res := transport.UserResponse{
		ID:       user.ID,
		Nickname: user.Nickname,
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error encoding users response.\n", http.StatusInternalServerError)
	}
}
