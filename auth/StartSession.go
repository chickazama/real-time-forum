package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid/v5"
)

func StartSession(w http.ResponseWriter, userID int) error {
	_, exists := sessionStore[userID]
	if exists {
		return errors.New("session already exists")
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	value := uuid.String()
	idCookie := http.Cookie{
		Name:     idCookieName,
		Value:    fmt.Sprintf("%d", userID),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   0,
		SameSite: http.SameSiteLaxMode,
	}
	sessionCookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   0,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &idCookie)
	http.SetCookie(w, &sessionCookie)
	sessionStore[userID] = value
	return nil
}
