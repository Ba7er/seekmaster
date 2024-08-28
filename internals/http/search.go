package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query("SELECT first_name FROM customers")
	if err != nil {
		http.Error(w, "Failed to query database", http.StatusInternalServerError)
		log.Printf("Error querying database: %v", err)
		return
	}
	defer rows.Close()

	// Slice to store customer names.
	var people []string

	// Iterate through the result rows.
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			log.Printf("Error scanning row: %v", err)
			return
		}
		people = append(people, name)
	}

	// Check for errors during the rows iteration.
	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating rows", http.StatusInternalServerError)
		log.Printf("Error iterating rows: %v", err)
		return
	}

	// Encode the people slice as JSON and send it in the response.
	if err := json.NewEncoder(w).Encode(people); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
		return
	}
}
