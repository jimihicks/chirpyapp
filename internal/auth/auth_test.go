package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckPasswordHash(t *testing.T) {
	// First, we need to create some hashed passwords for testing
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJWT(t *testing.T) {
	t.Run("round-trip valid", func(t *testing.T) {
		id := uuid.New()
		secret := "secret"
		tok, err := MakeJWT(id, secret, time.Minute)
		if err != nil || tok == "" {
			t.Fatalf("MakeJWT: %v", err)
		}
		got, err := ValidateJWT(tok, secret)
		if err != nil {
			t.Fatalf("ValidateJWT: %v", err)
		}
		if got != id {
			t.Fatalf("want %s got %s", id, got)
		}
	})

	t.Run("expired", func(t *testing.T) {
		id := uuid.New()
		secret := "secret"
		tok, err := MakeJWT(id, secret, -time.Minute)
		if err != nil || tok == "" {
			t.Fatalf("MakeJWT: %v", err)
		}
		if _, err := ValidateJWT(tok, secret); err == nil {
			t.Fatal("expected error for expired token")
		}
	})

	t.Run("wrong secret", func(t *testing.T) {
		id := uuid.New()
		tok, err := MakeJWT(id, "a", time.Minute)
		if err != nil || tok == "" {
			t.Fatalf("MakeJWT: %v", err)
		}
		if _, err := ValidateJWT(tok, "b"); err == nil {
			t.Fatal("expected error for wrong secret")
		}
	})
}
