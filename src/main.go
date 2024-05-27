package main

import (
	"net/http"
	"salesList/src/handlers"
)

func main() {
	// Map endpoints to handler functions
	http.HandleFunc("/", handlers.RootHandler)
	// Add more mappings for other endpoints if needed

	// Start the HTTP server on port 8080
	http.ListenAndServe("localhost:8080", nil)
}
