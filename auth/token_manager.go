package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"sync"
)

// Token struct keeps track of how many uses are left for a token
type Token struct {
	UsesLeft int
}

// tokenStore holds all active tokens and their usage state
var tokenStore = make(map[string]*Token)
var mu sync.Mutex // Mutex to safely access tokenStore concurrently

// generateToken creates a cryptographically secure 64-character token
func generateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Error generating secure token:", err)
	}
	return hex.EncodeToString(b)
}

// CreateToken generates and stores a new token with 5 uses
func CreateToken() string {
	mu.Lock()
	defer mu.Unlock()

	t := generateToken()
	tokenStore[t] = &Token{UsesLeft: 5}
	log.Printf("✅ Token generated: %s (5 uses)", t)
	return t
}

// ValidateToken checks if a token is valid and still has remaining uses.
// If valid, it decrements the use count. Otherwise, it logs and returns an error.
func ValidateToken(t string) (bool, string) {
	mu.Lock()
	defer mu.Unlock()

	token, exists := tokenStore[t]
	if !exists {
		log.Printf("❌ Invalid token received: %s", t)
		return false, "Invalid token"
	}
	if token.UsesLeft <= 0 {
		log.Printf("⚠️ Expired token used: %s", t)
		return false, "Token expired"
	}

	token.UsesLeft--
	log.Printf("🔁 Token used: %s (remaining uses: %d)", t, token.UsesLeft)
	return true, ""
}
