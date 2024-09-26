package main

import (
	"github.com/Ba7er/seekmaster/internals/http"
)

func run() {

	server := http.NewServer()
	server.InitDatabaseConnection()
	server.LoadRouters()
	server.Open()

}
func main() {
	run()
}
