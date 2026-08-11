package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// newTestHandler builds a Handler wired to the test database.
func newTestHandler(t *testing.T) (*Handler, *sql.DB) {
	t.Helper()

	db := testDB(t)
	h := NewHandler(NewService(NewRepository(db)))
	return h, db
}

// registerBody builds a valid registration body with unique values.
func registerBody() string {
	u := uniqueUser()
	return fmt.Sprintf(
		`{"name":%q,"username":%q,"email":%q,"password":"password123"}`,
		u.Name, u.Username, u.Email,
	)
}

// performRegister sends a POST /auth/register request and returns the
// recorded response.
func performRegister(h *Handler, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.Register(rec, req)
	return rec
}

func TestRegisterHandlerValid(t *testing.T) {
	h, db := newTestHandler(t)

	rec := performRegister(h, registerBody())

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d: %s", rec.Code, rec.Body.String())
	}

	var resp RegisterResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	if resp.User.ID == "" {
		t.Fatal("expected the response to include the user ID")
	}
	if resp.User.Email == "" {
		t.Fatal("expected the response to include the user email")
	}
	if strings.Contains(rec.Body.String(), "password_hash") {
		t.Fatal("response must never include the password hash")
	}

	removeUser(t, db, resp.User.ID)
}

func TestRegisterHandlerMissingRequiredField(t *testing.T) {
	h, _ := newTestHandler(t)

	rec := performRegister(h, `{"name":"Test User","username":"tester123","password":"password123"}`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestRegisterHandlerInvalidEmail(t *testing.T) {
	h, _ := newTestHandler(t)

	rec := performRegister(h, `{"name":"Test User","username":"tester123","email":"not-an-email","password":"password123"}`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestRegisterHandlerShortPassword(t *testing.T) {
	h, _ := newTestHandler(t)

	rec := performRegister(h, `{"name":"Test User","username":"tester123","email":"short@example.com","password":"short"}`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestRegisterHandlerDuplicateEmail(t *testing.T) {
	h, db := newTestHandler(t)

	body := registerBody()

	first := performRegister(h, body)
	if first.Code != http.StatusCreated {
		t.Fatalf("first registration: expected status 201, got %d", first.Code)
	}

	var resp RegisterResponse
	if err := json.Unmarshal(first.Body.Bytes(), &resp); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	defer removeUser(t, db, resp.User.ID)

	second := performRegister(h, body)
	if second.Code != http.StatusConflict {
		t.Fatalf("expected status 409 for duplicate email, got %d: %s", second.Code, second.Body.String())
	}
}

func TestRegisterHandlerDuplicateUsername(t *testing.T) {
	h, db := newTestHandler(t)

	u := uniqueUser()
	firstBody := fmt.Sprintf(
		`{"name":%q,"username":%q,"email":%q,"password":"password123"}`,
		u.Name, u.Username, u.Email,
	)

	first := performRegister(h, firstBody)
	if first.Code != http.StatusCreated {
		t.Fatalf("first registration: expected status 201, got %d", first.Code)
	}

	var resp RegisterResponse
	if err := json.Unmarshal(first.Body.Bytes(), &resp); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	defer removeUser(t, db, resp.User.ID)

	// Same username, different email.
	other := uniqueUser()
	secondBody := fmt.Sprintf(
		`{"name":%q,"username":%q,"email":%q,"password":"password123"}`,
		u.Name, u.Username, other.Email,
	)

	second := performRegister(h, secondBody)
	if second.Code != http.StatusConflict {
		t.Fatalf("expected status 409 for duplicate username, got %d: %s", second.Code, second.Body.String())
	}
}

func TestRegisterHandlerInvalidJSON(t *testing.T) {
	h, _ := newTestHandler(t)

	rec := performRegister(h, "{this is not json")

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestRegisterHandlerWrongMethod(t *testing.T) {
	h, _ := newTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/auth/register", nil)
	rec := httptest.NewRecorder()
	h.Register(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rec.Code)
	}
}
