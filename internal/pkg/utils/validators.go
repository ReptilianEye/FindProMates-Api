package utils

import (
	"fmt"
	"net/mail"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
func ValidateName(name string) (string, error) {
	if len(name) < 3 {
		return "", fmt.Errorf("name must be at least 3 characters long")
	}
	return cases.Title(language.English).String(name), nil
}
