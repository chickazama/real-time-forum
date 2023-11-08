package ui

import (
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/repo"
	"net/http"
)

type LogoutHandler struct {
	Repo repo.IRepository
}

func NewLogoutHandler(r repo.IRepository) *LogoutHandler {
	return &LogoutHandler{Repo: r}
}

func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed.\n", http.StatusMethodNotAllowed)
		return
	}
	userID, err := auth.AuthorizeRequest(r)
	if err != nil {
		log.Println(err.Error())
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
