package main

import (
	"github/Ba7er/seekmaster/internals/http"
)

// func connectDb() {
// 	var cfg = mysql.Config{
// 		User:   "myuser",
// 		Passwd: "mypassword",
// 		Net:    "tcp",
// 		Addr:   "localhost:3306",
// 		DBName: "myapp",
// 	}

// 	db, err := sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr, "Could Not Connect to DB")
// 		os.Exit(1)
// 	}
// 	log.Print("Connected")
// }
//connectDb()

func run() {
	server := http.NewServer()
	server.Open()

}
func main() {
	run()
}
