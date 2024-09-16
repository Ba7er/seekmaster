package services

// import (
// 	"database/sql"
// 	"log"
// 	"net/http"
// )

// type SearchQueryParams struct {
// 	q           string
// 	brand       string
// 	category    string
// 	subCategory string
// 	sort        string
// 	limit       int
// }

// type Product struct {
// 	Name          string
// 	Description   string
// 	Price         float64
// 	Stock         int
// 	CategoryID    int
// 	SubCategoryID int
// 	BrandID       int
// }

// func Search(params SearchQueryParams, db *sql.DB) (Product, error) {
// 	rows, err := db.Query("SELECT first_name FROM product")
// 	if err != nil {
// 		http.Error(w, "Failed to query database", http.StatusInternalServerError)
// 		log.Printf("Error querying database: %v", err)
// 		return
// 	}
// 	defer rows.Close()

// 	// Slice to store customer names.
// 	var people []string

// 	// Iterate through the result rows.
// 	for rows.Next() {
// 		var name string
// 		if err := rows.Scan(&name); err != nil {
// 			http.Error(w, "Error scanning row", http.StatusInternalServerError)
// 			log.Printf("Error scanning row: %v", err)
// 			return
// 		}
// 		people = append(people, name)
// 	}

// 	// Check for errors during the rows iteration.
// 	if err := rows.Err(); err != nil {
// 		http.Error(w, "Error iterating rows", http.StatusInternalServerError)
// 		log.Printf("Error iterating rows: %v", err)
// 		return
// 	}

// }
