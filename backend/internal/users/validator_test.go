package users

import (
	"strings"
	"testing"
)

func validUser() User {
	return User{
		Username:     "wanjiru",
		Email:        "wanjiru@example.com",
		PasswordHash: "hashed-password",
	}
}

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name:    "valid user",
			user:    validUser(),
			wantErr: false,
		},
		{
			name: "username too short",
			user: func() User {
				u := validUser()
				u.Username = "ab"
				return u
			}(),
			wantErr: true,
		},
		{
			name: "username too long",
			user: func() User {
				u := validUser()
				u.Username = strings.Repeat("a", 51)
				return u
			}(),
			wantErr: true,
		},
		{
			name: "empty email",
			user: func() User {
				u := validUser()
				u.Email = ""
				return u
			}(),
			wantErr: true,
		},
		{
			name: "invalid email",
			user: func() User {
				u := validUser()
				u.Email = "not-an-email"
				return u
			}(),
			wantErr: true,
		},
		{
			name: "empty password hash",
			user: func() User {
				u := validUser()
				u.PasswordHash = ""
				return u
			}(),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()

			if (err != nil) != tt.wantErr {
				t.Fatalf("Validate() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserIsValid(t *testing.T) {
	user := validUser()

	if !user.IsValid() {
		t.Fatal("expected valid user to return true")
	}

	user.Username = "ab"

	if user.IsValid() {
		t.Fatal("expected invalid user to return false")
	}
}
