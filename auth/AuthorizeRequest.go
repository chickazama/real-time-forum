package auth

import (
	"errors"
	"net/http"
	"strconv"
)

func AuthorizeRequest(r *http.Request) (int, error) {
	idCookie, err := r.Cookie("UserID")
	if err != nil {
		return -1, errors.New("user ID Cookie")
	}
	userID, err := strconv.Atoi(idCookie.Value)
	if err != nil {
		return -1, errors.New("user ID Cookie")
	}
	sessionCookie, err := r.Cookie("Session")
	if err != nil {
		return -1, errors.New("session Cookie")
	}
	if !validateSessionCookie(userID, sessionCookie.Value) {
		return -1, errors.New("session Cookie")
	}
	return userID, nil
}
