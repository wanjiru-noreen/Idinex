package auth

import "testing"

func TestValidateRegisterRequest(t *testing.T) {
	tests := []struct {
		name    string
		request RegisterRequest
		wantErr bool
	}{
		{
			name: "valid request",
			request: RegisterRequest{
				Email:    "user@example.com",
				Password: "password123",
				FullName: "Wanjiru Noreen",
			},
			wantErr: false,
		},
		{
			name: "missing full name",
			request: RegisterRequest{
				Email:    "user@example.com",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "invalid email",
			request: RegisterRequest{
				Email:    "invalid-email",
				Password: "password123",
				FullName: "Wanjiru Noreen",
			},
			wantErr: true,
		},
		{
			name: "short password",
			request: RegisterRequest{
				Email:    "user@example.com",
				Password: "pass1",
				FullName: "Wanjiru Noreen",
			},
			wantErr: true,
		},
		{
			name: "password without number",
			request: RegisterRequest{
				Email:    "user@example.com",
				Password: "password",
				FullName: "Wanjiru Noreen",
			},
			wantErr: true,
		},
		{
			name: "password without letter",
			request: RegisterRequest{
				Email:    "user@example.com",
				Password: "12345678",
				FullName: "Wanjiru Noreen",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRegisterRequest(tt.request)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateRegisterRequest() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLoginRequest(t *testing.T) {
	tests := []struct {
		name    string
		request LoginRequest
		wantErr bool
	}{
		{
			name: "valid request",
			request: LoginRequest{
				Email:    "user@example.com",
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			request: LoginRequest{
				Email:    "invalid-email",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "missing password",
			request: LoginRequest{
				Email: "user@example.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateLoginRequest(tt.request)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateLoginRequest() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
