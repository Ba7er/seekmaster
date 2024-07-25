package server

import (
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request)
type Adapter func(handler) handler

func SetJSONHeader() Adapter {
	return func(h handler) handler {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			h(w, r)
		}
	}
}

func AuthenticateXAPIKey() Adapter {
	return func(h handler) handler {
		return func(w http.ResponseWriter, r *http.Request) {
			h(w, r)
		}
	}
}

func Adapt(h handler, adapters ...Adapter) handler {
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}
	return h
}
