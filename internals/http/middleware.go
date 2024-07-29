package http

import (
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)
type Adapter func(Handler) Handler

func SetJSONHeader() Adapter {
	return func(h Handler) Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			h(w, r)
		}
	}
}

// @TODO write the logic for checking X-API-KEY
func AuthenticateXAPIKey() Adapter {
	return func(h Handler) Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			h(w, r)
		}
	}
}

func Adapt(h Handler, adapters ...Adapter) Handler {
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}
	return h
}
