package user

import (
	"errors"
	"regexp"
)
type Email struct {
	Address string
}

func NewEmail(address string) (Email, error) {
	if !regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`).MatchString(address) {
		return Email{}, ErrInvalidEmail
	}

	return Email{Address: address}, nil
}

var ErrInvalidEmail = errors.New("invalid email format")