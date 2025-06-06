package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file if available
	godotenv.Load()

	// Set up HTTP handlers for token generation and data fetching
	http.HandleFunc("/token", GenerateTokenHandler)
	http.HandleFunc("/characters", ValidateAndFetchHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})
	// Get port from environment or default to 8080
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Auth service running on port", port)
	// Start HTTP server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
