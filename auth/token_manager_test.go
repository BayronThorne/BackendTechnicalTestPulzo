package main

import (
	"testing"
)

// TestCreateToken checks that a new token is created and immediately valid.
func TestCreateToken(t *testing.T) {
	token := CreateToken()
	if token == "" {
		t.Fatal("Expected a token, got an empty string")
	}

	valid, _ := ValidateToken(token)
	if !valid {
		t.Fatal("Expected token to be valid immediately after creation")
	}
}

// TestTokenUsageLimit validates that a token becomes invalid after 5 uses.
func TestTokenUsageLimit(t *testing.T) {
	token := CreateToken()

	// Use the token 5 times
	for i := 0; i < 5; i++ {
		valid, _ := ValidateToken(token)
		if !valid {
			t.Fatalf("Token should be valid on use #%d", i+1)
		}
	}

	// 6th use should fail (token is expired)
	valid, msg := ValidateToken(token)
	if valid {
		t.Fatal("Expected token to be expired after 5 uses")
	}
	if msg != "Token expired" {
		t.Fatalf("Expected 'Token expired', got: %s", msg)
	}
}

// TestInvalidToken checks validation logic for a completely unknown token.
func TestInvalidToken(t *testing.T) {
	valid, msg := ValidateToken("nonexistenttoken123")
	if valid {
		t.Fatal("Expected validation to fail for unknown token")
	}
	if msg != "Invalid token" {
		t.Fatalf("Expected 'Invalid token', got: %s", msg)
	}
}
