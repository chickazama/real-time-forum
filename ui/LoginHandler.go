package ui

import (
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/dal"
	"matthewhope/real-time-forum/repo"
	"matthewhope/real-time-forum/transport"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	Repo repo.IRepository
}

func NewLoginHandler(r repo.IRepository) *LoginHandler {
	return &LoginHandler{Repo: r}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed.\n", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
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
	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(dto.Password))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
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
