package http

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func (s *Server) InitDatabaseConnection() {
	var cfg = mysql.Config{
		User:   "myuser",
		Passwd: "mypassword",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "myapp",
	}

	var err error
	s.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := s.db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Print("Db is Connected")
}
