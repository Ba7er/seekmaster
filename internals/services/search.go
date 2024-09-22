package services

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
)

type SearchQueryParams struct {
	q           string
	brand       string
	category    string
	subCategory string
	sort        string
	limit       int
}

type Product struct {
	Name          string
	Description   string
	Price         float64
	Stock         int
	CategoryID    int
	SubCategoryID int
	BrandID       int
}

func Search(keywords url.Values, db *sql.DB) ([]string, error) {
	fmt.Println(keywords)
	rows, err := db.Query("SELECT nam FROM product")
	if err != nil {
		return nil, fmt.Errorf("could not execute the query, ther err")
	}
	defer rows.Close()

	// Slice to store customer names.
	var product []string

	// Iterate through the result rows.
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {

			log.Printf("Error scanning row: %v", err)
			return nil, fmt.Errorf("error scanning row")
		}
		product = append(product, name)
	}

	// Check for errors during the rows iteration.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error scanning row")
	}
	return product, nil
}
