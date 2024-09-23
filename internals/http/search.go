package http

import (
	"encoding/json"
	"net/http"

	"github.com/Ba7er/seekmaster/internals/services"
)

func (s *Server) SearchHandler(w http.ResponseWriter, r *http.Request) {
	keywords := r.URL.Query()

	products, err := services.Search(keywords, s.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
