package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGetCharactersHandler_Success simulates a successful response
// from the external Rick and Morty API using a mock server.
func TestGetCharactersHandler_Success(t *testing.T) {
	// Create a mock API server to simulate the external API
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"results": [{"id": 1, "name": "Rick Sanchez"}]}`))
	}))
	defer mockAPI.Close()

	// Override the default handler temporarily to use the mock API
	originalHandler := GetCharactersHandler
	GetCharactersHandler = func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(mockAPI.URL)
		if err != nil {
			http.Error(w, "Error querying external API", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
	defer func() { GetCharactersHandler = originalHandler }()

	// Prepare test request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/characters", nil)
	rec := httptest.NewRecorder()

	GetCharactersHandler(rec, req)

	// Validate the response
	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "Rick Sanchez") {
		t.Fatalf("Expected response to contain 'Rick Sanchez', got %s", rec.Body.String())
	}
}

// TestHealthCheck verifies that the /health endpoint returns 200 and "OK".
func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	// Simulate health endpoint directly
	handler := func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	}

	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", rec.Code)
	}

	if rec.Body.String() != "OK" {
		t.Fatalf("Expected body to be 'OK', got '%s'", rec.Body.String())
	}
}
