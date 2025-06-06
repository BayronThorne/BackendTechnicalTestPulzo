package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// GenerateTokenHandler handles POST requests to create a new token.
// Returns the token as a JSON response.
func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("⚠️ Invalid method %s on /token", r.Method)
		return
	}

	token := CreateToken()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// ValidateAndFetchHandler handles GET requests to /characters.
// It validates the token and fetches character data from the datafetcher service.
func ValidateAndFetchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("⚠️ Invalid method %s on /characters", r.Method)
		return
	}

	token := r.Header.Get("Authorization")

	// Check if token was sent
	if token == "" {
		http.Error(w, "Token required", http.StatusUnauthorized)
		log.Println("❌ No token provided in request")
		return
	}

	// Check if token has correct length (64 chars from crypto/rand)
	if len(token) != 64 {
		http.Error(w, "Invalid token format", http.StatusBadRequest)
		log.Printf("❌ Token with invalid format: %s", token)
		return
	}

	// Validate the token
	valid, msg := ValidateToken(token)
	if !valid {
		http.Error(w, msg, http.StatusUnauthorized)
		log.Printf("❌ Token validation failed: %s", msg)
		return
	}

	// Prepare HTTP client with timeout
	client := &http.Client{Timeout: 5 * time.Second}
	log.Println("🌐 Calling datafetcher service...")

	resp, err := client.Get("http://datafetcher:8081/characters")
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		log.Printf("❌ Error calling datafetcher: %v", err)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	log.Println("✅ Characters fetched successfully from datafetcher")
}
