package main

import (
	"database/sql"
	"fmt"
	"github/Ba7er/seekmaster/internals/http"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "myuser",
		Passwd: "mypassword",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "myapp",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	log.Fatal(http.Run())
}
