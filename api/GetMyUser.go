package api

import (
	"encoding/json"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/transport"
	"net/http"
)

func GetMyUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed.\n", http.StatusMethodNotAllowed)
		return
	}
	userID, err := auth.GetUserIDFromSessionCookie(r)
	if err != nil {
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	user, err := dal.GetUserByID(userID)
	if err != nil {
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	userResponse := transport.UserResponse{
		ID:       user.ID,
		Nickname: user.Nickname,
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(&userResponse)
	if err != nil {
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
}
