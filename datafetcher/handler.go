package main

import (
	"io"
	"net/http"
)

// GetCharactersHandler handles requests for Rick and Morty characters
// It proxies the request to the external public API and returns the result
var GetCharactersHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		http.Error(w, "Error querying external API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the API response directly to the client, preserving headers and status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
