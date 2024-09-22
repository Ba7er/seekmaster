package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ba7er/seekmaster/internals/services"
)

func (s *Server) SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Encode the people slice as JSON and send it in the response.
	keyword := r.URL.Query()
	res, err := services.Search(keyword, s.db)
	if err != nil {
		http.Error(w, "Failed to query database", http.StatusInternalServerError)
		log.Printf("Error querying database: %v", err)
		return
	}
	people := make(map[string]string)
	people["name"] = "Abed"

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
		return
	}
}
