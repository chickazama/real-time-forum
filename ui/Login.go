package ui

import (
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/transport"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed.\n", http.StatusMethodNotAllowed)
		return
	}
	_, err := auth.GetUserIDFromSessionCookie(r)
	if err == nil {
		http.Error(w, "session already exists.\n", http.StatusNotAcceptable)
		return
	}
	err = r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	// Build Signup Request DTO
	dto := transport.UserLoginRequest{
		Nickname: r.PostFormValue("nickname"),
		Password: r.PostFormValue("password"),
	}
	// Validate DTO
	err = dto.Validate()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid details.\n", http.StatusBadRequest)
		return
	}
	user, err := dal.GetUserByNickname(dto.Nickname)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid details.\n", http.StatusBadRequest)
		return
	}
	err = auth.StartSession(w, user.ID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error starting session.\n", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
