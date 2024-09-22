package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Server struct {
	server *http.Server
	router *http.ServeMux
	db     *sql.DB
}

func (s *Server) ConnectDB() {
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
		//@TODO: Should we use code 1 or ?
		os.Exit(1)
	}
	log.Print("Db is Connected")
}

type CustomHandler struct{}

func (CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	people := make(map[string]string)
	people["name"] = "Abed"

	if err := json.NewEncoder(w).Encode(people); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Printf("Error encoding JSON: %v", err)
		return
	}
}
func (s *Server) LoadRouters() {

	s.router.Handle("GET /search", SetJSONHeader(http.HandlerFunc(s.SearchHandler)))
	s.server.Handler = s.router

}

func (s *Server) Open() {
	// @TODO: read about error handling when starting a server.
	s.server.ListenAndServe()
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{
			Addr:           ":9100",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		router: http.NewServeMux(),
	}

	return s
}
