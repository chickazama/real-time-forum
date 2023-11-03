package api

import (
	"encoding/json"
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/dal"
	"net/http"
	"strconv"
)

func GetChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed.\n", http.StatusMethodNotAllowed)
		return
	}
	idCookie, err := r.Cookie("UserID")
	if err != nil {
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	userID, err := strconv.Atoi(idCookie.Value)
	if err != nil {
		log.Println("here")
		log.Println(err.Error())
		http.Error(w, "bad request.\n", http.StatusBadRequest)
		return
	}
	sessionCookie, err := r.Cookie("Session")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	if !auth.ValidateSessionCookie(userID, sessionCookie.Value) {
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
	messages, err := dal.GetMessagesBySenderAndTargetIDs(userID, targetID)
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
