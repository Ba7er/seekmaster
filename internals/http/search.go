package http

import (
	"encoding/json"
	"net/http"

	"github.com/Ba7er/seekmaster/internals/services"
)

func (s *Server) SearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	products, err := services.Search(q, s.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Success bool               `json:"success"`
		Data    []services.Product `json:"data"`
	}{
		Success: true,
		Data:    products,
	}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
