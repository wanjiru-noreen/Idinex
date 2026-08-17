package users

import (
	"fmt"
	"net/mail"
	"strings"
)

func (u User) Validate() error {
	username := strings.TrimSpace(u.Username)
	email := strings.TrimSpace(u.Email)

	if len(username) < 3 {
		return fmt.Errorf("username must be at least 3 characters")
	}

	if len(username) > 50 {
		return fmt.Errorf("username must not exceed 50 characters")
	}

	if email == "" {
		return fmt.Errorf("email is required")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email")
	}

	if strings.TrimSpace(u.PasswordHash) == "" {
		return fmt.Errorf("password hash is required")
	}

	return nil
}
