package auth

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
)

func StartSession(w http.ResponseWriter, userID int) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	value := uuid.String()
	cookie := http.Cookie{
		Name:     sessionCookieName,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   0,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	sessionStore[value] = userID
	return nil
}
