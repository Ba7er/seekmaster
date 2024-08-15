package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
	mux *http.ServeMux
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	h.mux.ServeHTTP(w, r)
}

func (h *MyHandler) RegisterHandlers() {
	h.mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here")
		err := json.NewEncoder(w).Encode(&struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			ID:   2,
			Name: "Abed Dandashi with docker and mysql3",
		})
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	})
}

func main() {
	customHandler := MyHandler{
		mux: http.NewServeMux(),
	}
	customHandler.RegisterHandlers()
	s := &http.Server{
		Addr:           ":9100",
		Handler:        &customHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
