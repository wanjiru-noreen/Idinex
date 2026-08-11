package auth

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// testDB opens a connection to the database used by the tests.
// It looks for TEST_DATABASE_URL in the environment and falls back
// to the local development database.
func testDB(t *testing.T) *sql.DB {
	t.Helper()

	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/idinex?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		t.Fatalf("ping database: %v", err)
	}

	t.Cleanup(func() {
		_ = db.Close()
	})

	return db
}

// uniqueUser builds a user with a unique username and email so tests
// can run again and again without colliding with previous rows.
func uniqueUser() *User {
	now := time.Now().UnixNano()
	return &User{
		Name:         "Test User",
		Username:     fmt.Sprintf("test_user_%d", now),
		Email:        fmt.Sprintf("test_%d@example.com", now),
		PasswordHash: "a-non-empty-dummy-hash",
	}
}

// removeUser deletes the row created by a test so the database stays clean.
func removeUser(t *testing.T, db *sql.DB, id string) {
	t.Helper()

	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		t.Fatalf("cleanup: delete user: %v", err)
	}
}

func TestRepositoryCreateUser(t *testing.T) {
	db := testDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	user := uniqueUser()

	if err := repo.CreateUser(ctx, user); err != nil {
		t.Fatalf("CreateUser returned an error: %v", err)
	}

	if user.ID == "" {
		t.Fatal("expected the database to generate an ID")
	}
	if user.CreatedAt.IsZero() {
		t.Fatal("expected created_at to be set")
	}
	if user.UpdatedAt.IsZero() {
		t.Fatal("expected updated_at to be set")
	}

	removeUser(t, db, user.ID)
}

func TestRepositoryFindUserByEmail(t *testing.T) {
	db := testDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	user := uniqueUser()
	if err := repo.CreateUser(ctx, user); err != nil {
		t.Fatalf("CreateUser returned an error: %v", err)
	}
	defer removeUser(t, db, user.ID)

	found, err := repo.FindUserByEmail(ctx, user.Email)
	if err != nil {
		t.Fatalf("FindUserByEmail returned an error: %v", err)
	}
	if found.ID != user.ID {
		t.Fatalf("expected user ID %q, got %q", user.ID, found.ID)
	}
	if found.Name != user.Name {
		t.Fatalf("expected name %q, got %q", user.Name, found.Name)
	}
	if found.Username != user.Username {
		t.Fatalf("expected username %q, got %q", user.Username, found.Username)
	}
	if found.PasswordHash != user.PasswordHash {
		t.Fatal("expected password hash to be read back")
	}
}

func TestRepositoryFindUserByEmailNotFound(t *testing.T) {
	db := testDB(t)
	repo := NewRepository(db)

	_, err := repo.FindUserByEmail(context.Background(), "nobody@example.com")
	if err == nil {
		t.Fatal("expected an error for an unknown email")
	}
	if err != ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestRepositoryFindUserByID(t *testing.T) {
	db := testDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	user := uniqueUser()
	if err := repo.CreateUser(ctx, user); err != nil {
		t.Fatalf("CreateUser returned an error: %v", err)
	}
	defer removeUser(t, db, user.ID)

	found, err := repo.FindUserByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("FindUserByID returned an error: %v", err)
	}
	if found.ID != user.ID {
		t.Fatalf("expected user ID %q, got %q", user.ID, found.ID)
	}
	if found.Email != user.Email {
		t.Fatalf("expected email %q, got %q", user.Email, found.Email)
	}
}

func TestRepositoryFindUserByIDNotFound(t *testing.T) {
	db := testDB(t)
	repo := NewRepository(db)

	_, err := repo.FindUserByID(context.Background(), "00000000-0000-0000-0000-000000000000")
	if err == nil {
		t.Fatal("expected an error for an unknown ID")
	}
	if err != ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestRepositoryCreateUserDuplicateEmail(t *testing.T) {
	db := testDB(t)
	repo := NewRepository(db)
	ctx := context.Background()

	user := uniqueUser()
	if err := repo.CreateUser(ctx, user); err != nil {
		t.Fatalf("CreateUser returned an error: %v", err)
	}
	defer removeUser(t, db, user.ID)

	duplicate := uniqueUser()
	duplicate.Email = user.Email

	if err := repo.CreateUser(ctx, duplicate); err == nil {
		t.Fatal("expected an error when creating a user with a duplicate email")
	}
}
