package ui

import (
	"log"
	"matthewhope/real-time-forum/auth"
	"matthewhope/real-time-forum/repo"
	"matthewhope/real-time-forum/transport"
	"net/http"
)

type SignupHandler struct {
	Repo repo.IRepository
}

func NewSignupHandler(r repo.IRepository) *SignupHandler {
	return &SignupHandler{Repo: r}
}

func (h *SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	dto := transport.UserSignupRequest{
		Nickname:  r.PostFormValue("nickname"),
		Age:       r.PostFormValue("age"),
		Gender:    r.PostFormValue("gender"),
		FirstName: r.PostFormValue("firstName"),
		LastName:  r.PostFormValue("lastName"),
		Email:     r.PostFormValue("emailAddress"),
		Password:  r.PostFormValue("password"),
	}
	// Validate DTO
	err = dto.Validate()
	if err != nil {
		http.Error(w, "invalid details.\n", http.StatusBadRequest)
		return
	}
	// Add User to DB
	userID, err := h.Repo.CreateUser(dto.Nickname, dto.Age, dto.Gender, dto.FirstName, dto.LastName, dto.Email, dto.Password)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "user with details already exists.\n", http.StatusNotAcceptable)
		return
	}
	err = auth.StartSession(w, userID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "error starting session.\n", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
