package auth

import "time"

// RegisterRequest is the JSON body the client sends to POST /auth/register.
type RegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserDTO is the safe representation of a user sent to the client.
// The password hash is never included.
type UserDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// RegisterResponse is the JSON body returned after registration.
type RegisterResponse struct {
	Message string  `json:"message"`
	User    UserDTO `json:"user"`
}
