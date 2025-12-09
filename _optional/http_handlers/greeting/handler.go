// Package greeting provides an HTTP endpoint for the greeting function.
// This is an example of how to expose worker functionality via HTTP for frontend use.
package greeting

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Input defines the request body structure
type Input struct {
	Name string `json:"name"`
}

// Output defines the response body structure
type Output struct {
	Message string `json:"message"`
}

// Register registers the greeting HTTP endpoint.
// Call this with your router's mux to expose POST /api/greeting
func Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/greeting", handleGreeting)
}

// handleGreeting processes greeting requests
func handleGreeting(w http.ResponseWriter, r *http.Request) {
	// Set JSON content type
	w.Header().Set("Content-Type", "application/json")

	// Parse request body
	var input Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON: " + err.Error(),
		})
		return
	}

	// Validate input
	if input.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "name is required",
		})
		return
	}

	// Generate response (mirrors the worker function logic)
	output := Output{
		Message: fmt.Sprintf("Hello, %s!", input.Name),
	}

	// Send response
	json.NewEncoder(w).Encode(output)
}

