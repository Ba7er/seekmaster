package server

import (
	"encoding/json"
	"net/http"
)

type route struct {
	method  string
	path    string
	handler handler
}

var routes = []route{
	{"GET", "/search", search},
}

func Run() *http.ServeMux {
	mux := http.NewServeMux()
	for _, v := range routes {
		mux.HandleFunc(v.method+" "+v.path, Adapt(v.handler,
			SetJSONHeader(),
			AuthenticateXAPIKey(),
		))
	}
	return mux
}

func search(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   2,
		Name: "Abed Dandashi",
	})
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
