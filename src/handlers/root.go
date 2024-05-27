package handlers

import (
	"encoding/json"
	"net/http"
)

// Define a struct to represent your JSON data
type Response struct {
	Message string `json:"message"`
}

// Handler for the root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	// Create a response object
	response := Response{
		Message: "Hello, world!",
	}

	// Convert the response object to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonResponse)
}
