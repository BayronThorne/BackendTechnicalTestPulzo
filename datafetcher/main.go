package main

import (
	"log"
	"net/http"
)

func main() {
	// Register the handler for the /characters endpoint
	http.HandleFunc("/characters", GetCharactersHandler)

	// Log a message to indicate the service is running
	log.Println("Data fetcher running on port 8081")

	// Start the HTTP server on port 8081 and log any fatal errors
	log.Fatal(http.ListenAndServe(":8081", nil))
}
