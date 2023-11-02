package transport

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	inputMinLength = 2
	inputMaxLength = 50
)

type UserSignupRequest struct {
	Nickname  string
	Age       string
	Gender    string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (dto UserSignupRequest) Validate() error {
	if len(dto.Nickname) < inputMinLength || len(dto.Nickname) > inputMaxLength {
		return errors.New(INVALID_INPUT_NICK_NAME)
	}
	if len(dto.FirstName) < inputMinLength || len(dto.FirstName) > inputMaxLength {
		fmt.Println("invalid first name length")
		return errors.New(INVALID_INPUT_FIRST_NAME)
	}
	if len(dto.LastName) < inputMinLength || len(dto.LastName) > inputMaxLength {
		return errors.New(INVALID_INPUT_LAST_NAME)
	}
	emailPattern := `^[^@\s]+@[^@\s]+.[^@\s]`
	match, err := regexp.MatchString(emailPattern, dto.Email)
	if err != nil || !match {
		return errors.New(INVALID_INPUT_EMAIL)
	}
	return nil
}
