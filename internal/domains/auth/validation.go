package auth

import (
	"net/mail"
	"strings"
)

func validateRegisterRequest(email string, password string, name string) error {
	if email == "" {
		return ErrEmailRequired
	}

	if !isValidEmail(email) {
		return ErrInvalidEmail
	}

	if password == "" {
		return ErrPasswordRequired
	}

	if len(password) < 8 {
		return ErrPasswordTooShort
	}

	if len(password) > 128 {
		return ErrPasswordTooLong
	}

	name = strings.TrimSpace(name)

	if name == "" {
		return ErrNameRequired
	}

	if len(name) > 255 {
		return ErrNameTooLong
	}

	return nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}
