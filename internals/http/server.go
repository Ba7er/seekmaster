package http

import (
	"fmt"
	search "github/Ba7er/seekmaster/internals/services"
	"net/http"
	"time"
)

type route struct {
	method  string
	path    string
	handler Handler
}

var routes = []route{
	{"GET", "/search", Search},
}

type Server struct {
	serachService search.SearchService
	server        *http.Server
	router        *http.ServeMux
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

	for _, r := range routes {
		s.router.HandleFunc(r.method+" "+r.path, Adapt(r.handler,
			SetJSONHeader(),
			AuthenticateXAPIKey(),
		))
	}
	res := s.serachService.SearchForSomething()
	fmt.Println(res)

	return s
}

func (s *Server) Open() {
	s.server.Handler = s.router
	// @TODO: read about error handling when starting a server.
	s.server.ListenAndServe()

}
