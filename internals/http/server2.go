package http

import (
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

type CustomHandler struct {
	mux *http.ServeMux
}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *CustomHandler) registerHandlers() {
	for _, r := range routes {
		h.mux.HandleFunc(r.method+" "+r.path, Adapt(r.handler,
			SetJSONHeader(),
			AuthenticateXAPIKey(),
		))
	}
}

type Server struct {
	Server *http.Server
}

func NewServer() *Server {
	customHandler := CustomHandler{
		mux: http.NewServeMux(),
	}
	customHandler.registerHandlers()
	server := &Server{
		Server: &http.Server{
			Addr:           ":9100",
			Handler:        &customHandler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
	return server
}
