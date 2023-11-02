package auth

import (
	"errors"
	"net/http"
)

func GetUserIDFromSessionCookie(r *http.Request) (int, error) {
	var result int
	cookiePtr, err := r.Cookie(sessionCookieName)
	if err != nil {
		return result, err
	}
	result, exists := sessionStore[cookiePtr.Value]
	if !exists {
		return result, errors.New("no session")
	}
	return result, nil
}
