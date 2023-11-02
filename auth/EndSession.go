package auth

import (
	"errors"
	"net/http"
	"time"
)

func EndSession(w http.ResponseWriter, r *http.Request) error {
	cookiePtr, err := r.Cookie(sessionCookieName)
	if err != nil {
		return err
	}
	value := cookiePtr.Value
	_, exists := sessionStore[value]
	if !exists {
		return errors.New("session does not exist")
	}
	delete(sessionStore, value)
	cookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	return nil
}
