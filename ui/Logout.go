package ui

import (
	"matthewhope/real-time-forum/auth"
	"net/http"
	"strconv"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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
		http.Error(w, "bad request.\n", http.StatusBadRequest)
		return
	}
	sessionCookie, err := r.Cookie("Session")
	if err != nil {
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	if !auth.ValidateSessionCookie(userID, sessionCookie.Value) {
		http.Error(w, "unauthorized.\n", http.StatusUnauthorized)
		return
	}
	err = auth.EndSession(w, r, userID)
	if err != nil {
		http.Error(w, "internal server error.\n", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
