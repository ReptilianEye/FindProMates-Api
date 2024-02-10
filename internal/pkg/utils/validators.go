package utils

import (
	"fmt"
	"net/mail"
	"strings"
)

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
func ValidateName(name string) (*string, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("name must be at least 3 characters long")
	}
	name = strings.ToTitle(name)
	return &name, nil
}
