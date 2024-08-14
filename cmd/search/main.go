package main

import (
	"github/Ba7er/seekmaster/internals/http"
)

func main() {

	server := http.NewServer()
	server.Server.ListenAndServe()
}
