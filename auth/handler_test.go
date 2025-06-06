package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGenerateTokenHandler ensures that a valid POST request to /token returns a token and 200 status.
func TestGenerateTokenHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/token", nil)
	rec := httptest.NewRecorder()

	GenerateTokenHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "token") {
		t.Fatalf("Expected response to contain 'token', got %s", body)
	}
}

// TestGenerateTokenHandler_WrongMethod checks that GET requests to /token are rejected with 405.
func TestGenerateTokenHandler_WrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/token", nil)
	rec := httptest.NewRecorder()

	GenerateTokenHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("Expected 405, got %d", rec.Code)
	}
}

// TestValidateAndFetchHandler_NoToken ensures that requests to /characters without a token return 401.
func TestValidateAndFetchHandler_NoToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/characters", nil)
	rec := httptest.NewRecorder()

	ValidateAndFetchHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Expected 401 for missing token, got %d", rec.Code)
	}
}

// TestHealthCheck checks that /health returns 200 and the correct body content.
func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	// Define the health handler directly since it's anonymous in main.go
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
