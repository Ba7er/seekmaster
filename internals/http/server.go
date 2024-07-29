package http

import (
	"errors"
	"fmt"
	"github/Ba7er/seekmaster/internals/handlers"
	"net/http"
)

type route struct {
	method  string
	path    string
	handler Handler
}

var routes = []route{
	{"GET", "/search", handlers.Search},
}

func Run() error {
	mux := http.NewServeMux()
	for _, v := range routes {
		mux.HandleFunc(v.method+" "+v.path, Adapt(v.handler,
			SetJSONHeader(),
			AuthenticateXAPIKey(),
		))
	}
	err := http.ListenAndServe(":9100", mux)
	// to be investigated
	if !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("%s", err)
	}
	return nil
}
