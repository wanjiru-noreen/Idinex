package auth

import (
	"fmt"
	"net/mail"
	"strings"
	"unicode"
)

func ValidateRegisterRequest(req RegisterRequest) error {
	if strings.TrimSpace(req.FullName) == "" {
		return fmt.Errorf("full name is required")
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return fmt.Errorf("invalid email")
	}

	if len(req.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	var hasLetter, hasNumber bool

	for _, char := range req.Password {
		if unicode.IsLetter(char) {
			hasLetter = true
		}
		if unicode.IsNumber(char) {
			hasNumber = true
		}
	}

	if !hasLetter || !hasNumber {
		return fmt.Errorf("password must contain at least one letter and one number")
	}

	return nil
}

func ValidateLoginRequest(req LoginRequest) error {
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return fmt.Errorf("invalid email")
	}

	if strings.TrimSpace(req.Password) == "" {
		return fmt.Errorf("password is required")
	}

	return nil
}
