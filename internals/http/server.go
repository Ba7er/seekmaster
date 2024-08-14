package http

// import (
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/go-sql-driver/mysql"
// )

// // type route struct {
// // 	method  string
// // 	path    string
// // 	handler Handler
// // }

// var db *sql.DB

// // Capture connection properties.
// func connect() {
// 	var cfg = mysql.Config{
// 		User:   "myuser",
// 		Passwd: "mypassword",
// 		Net:    "tcp",
// 		Addr:   "localhost:3306",
// 		DBName: "myapp",
// 	}
// 	// Get a database handle.
// 	var err error

// 	db, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr, "Error in connecting to DB")
// 	}
// 	// rows, err := db.Query("SELECT * FROM customers")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// type customer struct {
// 	// 	firstName string
// 	// }
// 	// for rows.Next() {
// 	// 	var person customer
// 	// 	rows.Scan(&person.firstName)
// 	// 	fmt.Println(person)
// 	// }
// 	// defer rows.Close()
// 	log.Print("Connected")
// }

// // var routes = []route{
// // 	{"GET", "/search", handlers.Search},
// // }

// func Run() error {
// 	connect()
// 	mux := http.NewServeMux()
// 	for _, v := range routes {
// 		mux.HandleFunc(v.method+" "+v.path, Adapt(v.handler,
// 			SetJSONHeader(),
// 			AuthenticateXAPIKey(),
// 		))
// 	}
// 	err := http.ListenAndServe(":9100", mux)
// 	// to be investigated
// 	if !errors.Is(err, http.ErrServerClosed) {
// 		return fmt.Errorf("%s", err)
// 	}
// 	return nil
// }
