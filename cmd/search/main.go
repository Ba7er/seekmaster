package main

import (
	"github/Ba7er/seekmaster/internals/http"
)

func run() {

	server := http.NewServer()
	server.ConnectDB()
	server.LoadRouters()
	server.Open()

}
func main() {
	run()
}
