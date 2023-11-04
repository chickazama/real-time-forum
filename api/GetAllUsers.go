package api

import (
	"encoding/json"
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/transport"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed.\n", http.StatusMethodNotAllowed)
		return
	}
	_, err := auth.AuthorizeRequest(r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	users, err := dal.GetAllUsers()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	var usersResponse []transport.UserResponse
	for _, user := range users {
		userRes := transport.UserResponse{
			ID:       user.ID,
			Nickname: user.Nickname,
		}
		usersResponse = append(usersResponse, userRes)
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&usersResponse)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
}
