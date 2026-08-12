package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/lib/pq"
)

// ErrEmailTaken is returned when a user tries to register with an email
// that already has an account.
var ErrEmailTaken = errors.New("email is already registered")

// ErrUsernameTaken is returned when a user tries to register with a
// username that is already in use.
var ErrUsernameTaken = errors.New("username is already taken")

// Service contains the business rules for authentication.
type Service struct {
	repo *Repository
}

// NewService creates an authentication Service.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// RegisterUser creates a new account and returns safe user information.
// The password is hashed before it is stored.
func (s *Service) RegisterUser(ctx context.Context, req RegisterRequest) (UserDTO, error) {
	// Normalize values before storing them.
	req.Name = strings.TrimSpace(req.Name)
	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	// Reject an email that already has an account.
	_, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err == nil {
		return UserDTO{}, ErrEmailTaken
	}
	if !errors.Is(err, ErrNotFound) {
		return UserDTO{}, err
	}

	// Hash the password so the raw value is never stored.
	hash, err := HashPassword(req.Password)
	if err != nil {
		return UserDTO{}, err
	}

	user := &User{
		Name:         req.Name,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hash,
	}

	// Create the user. The unique constraints on email and username
	// protect against two registrations happening at the same time.
	if err := s.repo.CreateUser(ctx, user); err != nil {
		if isUniqueViolation(err) {
			return UserDTO{}, mapConstraintError(err)
		}
		return UserDTO{}, err
	}

	return UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

// isUniqueViolation reports whether the error comes from a UNIQUE
// constraint in the database.
func isUniqueViolation(err error) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code == "23505"
}

// mapConstraintError converts a database unique-constraint error into
// the friendliest message for the user.
func mapConstraintError(err error) error {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		if pqErr.Constraint == "users_username_unique" {
			return ErrUsernameTaken
		}
		if pqErr.Constraint == "users_email_unique" {
			return ErrEmailTaken
		}
	}
	return err
}
