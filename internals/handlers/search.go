package handlers

import (
	"encoding/json"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   2,
		Name: "Abed Dandashiiiii",
	})
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
