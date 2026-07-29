package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "mypassword"

	hash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hash == "" {
		t.Fatal("expected hash to be generated")
	}

	if hash == password {
		t.Fatal("password should not equal its hash")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "mypassword"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = CheckPassword(password, hash)
	if err != nil {
		t.Fatal("expected password to match hash")
	}

	err = CheckPassword("wrongpassword", hash)
	if err == nil {
		t.Fatal("expected password verification to fail")
	}
}
