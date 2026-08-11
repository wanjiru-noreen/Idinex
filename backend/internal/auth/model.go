package auth

import "time"

// User represents a registered application user.
// ID is stored as a string because the database generates the UUID
// and lib/pq can scan it into a string safely.
type User struct {
	ID           string
	Name         string
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
