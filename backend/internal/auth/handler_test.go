package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterHandlerInvalidRequest(t *testing.T) {
	body := `{
		"email": "invalid-email",
		"password": "password123",
		"full_name": "Wanjiru Noreen"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		strings.NewReader(body),
	)
	rec := httptest.NewRecorder()

	RegisterHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func TestRegisterHandlerValidRequest(t *testing.T) {
	body := `{
		"email": "user@example.com",
		"password": "password123",
		"full_name": "Wanjiru Noreen"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		strings.NewReader(body),
	)
	rec := httptest.NewRecorder()

	RegisterHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}
}

func TestRegisterHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/register", nil)
	rec := httptest.NewRecorder()

	RegisterHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rec.Code)
	}
}

func TestRegisterHandlerReturnsStandardizedError(t *testing.T) {
	body := `{
		"email": "invalid-email",
		"password": "password123",
		"full_name": "Wanjiru Noreen"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		strings.NewReader(body),
	)
	rec := httptest.NewRecorder()

	RegisterHandler(rec, req)

	var response struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}

	if response.Error != "invalid_request" {
		t.Fatalf(
			"expected error code %q, got %q",
			"invalid_request",
			response.Error,
		)
	}

	if response.Message == "" {
		t.Fatal("expected error message, got empty string")
	}
}
