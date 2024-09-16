package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Encode the people slice as JSON and send it in the response.
	people := make(map[string]string)
	people["name"] = "Abed"

	if err := json.NewEncoder(w).Encode(people); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
		return
	}
}
