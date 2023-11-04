package api

import (
	"encoding/json"
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/repo"
	"matthewhope/real-time-forum/transport"
	"net/http"
)

type UsersHandler struct {
	GetUsersRepo repo.GetUsersRepository
}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{GetUsersRepo: dal.NewDefaultGetUsersRepository()}
}

func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := auth.AuthorizeRequest(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	users, err := h.GetUsersRepo.GetUsers()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error retrieving users from database.\n", http.StatusInternalServerError)
		return
	}
	var res []transport.UserResponse
	for _, user := range users {
		dto := transport.UserResponse{
			ID:       user.ID,
			Nickname: user.Nickname,
		}
		res = append(res, dto)
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error encoding users response.\n", http.StatusInternalServerError)
	}
}
