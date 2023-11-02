package auth

import (
	"errors"
	"net/http"
	"time"
)

func EndSession(w http.ResponseWriter, r *http.Request, userID int) error {
	_, exists := sessionStore[userID]
	if !exists {
		return errors.New("session does not exist")
	}
	delete(sessionStore, userID)
	idCookie := http.Cookie{
		Name:     idCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteLaxMode,
	}
	sessionCookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &idCookie)
	http.SetCookie(w, &sessionCookie)
	return nil
}
