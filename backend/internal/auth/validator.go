package auth

import (
	"fmt"
	"regexp"
	"strings"
)

// Minimum accepted lengths for user-provided values.
const (
	MinUsernameLength = 3
	MinPasswordLength = 8
)

// emailPattern is a simple check that an email looks like name@domain.tld.
var emailPattern = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

// validateRegister checks a registration request and returns an error
// message describing the first problem, or an empty string if the
// request is valid.
func validateRegister(req RegisterRequest) string {
	if strings.TrimSpace(req.Name) == "" {
		return "name is required"
	}
	if len(strings.TrimSpace(req.Username)) < MinUsernameLength {
		return fmt.Sprintf("username must be at least %d characters", MinUsernameLength)
	}
	if !emailPattern.MatchString(strings.TrimSpace(req.Email)) {
		return "invalid email address"
	}
	if len(req.Password) < MinPasswordLength {
		return fmt.Sprintf("password must be at least %d characters", MinPasswordLength)
	}
	return ""
}
