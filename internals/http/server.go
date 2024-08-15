package http

import (
	"encoding/json"
	"github/Ba7er/seekmaster/internals/handlers"
	"net/http"
	"time"
)

type route struct {
	method  string
	path    string
	handler Handler
}

var routes = []route{
	{"GET", "/search", handlers.Search},
}

// type CustomHandler struct {
// 	mux *http.ServeMux
// }

// func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	h.mux.ServeHTTP(w, r)
// }

// func (h *CustomHandler) registerHandlers() {
// 	for _, r := range routes {
// 		h.mux.HanleFunc(r.method+" "+r.path, Adapt(r.handler,
// 			SetJSONHeader(),
// 			AuthenticateXAPIKey(),
// 		))
// 	}
// }

type Server struct {
	server *http.Server
	routes *http.ServeMux
}

func NewServer() *Server {
	//customHandler := CustomHandler{
	//mux: http.NewServeMux(),
	//}

	//customHandler.registerHandlers()

	server := &Server{

		server: &http.Server{
			Addr:           ":9100",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},

		routes: http.NewServeMux(),
	}

	// for _, r := range routes {
	server.routes.Handle("GET /hc", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
	//}
	return server
}

func (s *Server) Open() {
	s.server.Handler = s.routes
	s.server.ListenAndServe()

}
