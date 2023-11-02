package transport

import "errors"

type UserLoginRequest struct {
	Nickname string
	Password string
}

func (dto UserLoginRequest) Validate() error {
	// Login Policy Logic
	if len(dto.Nickname) < inputMinLength || len(dto.Nickname) > inputMaxLength {
		return errors.New(INVALID_INPUT_NICK_NAME)
	}
	if len(dto.Password) < inputMinLength || len(dto.Password) > inputMaxLength {
		return errors.New(INVALID_INPUT_PASSWORD)
	}
	return nil
}
