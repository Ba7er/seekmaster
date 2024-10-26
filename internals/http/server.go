package http

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	router *http.ServeMux
	db     *sql.DB
}

func (s *Server) LoadRouters() {
	s.router.Handle("GET /search", SetJSONHeader(http.HandlerFunc(s.SearchHandler)))
	s.server.Handler = s.router
}

func (s *Server) Start() {
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatalf("Couldn't start the server %s\n", err)
	}
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{
			Addr:           ":9100",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		router: http.NewServeMux(),
	}

	return s
}
