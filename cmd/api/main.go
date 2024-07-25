package main

import (
	"log"
	"net/http"
	"seekmaster/cmd/api/server"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", server.Run()))
}
