package services

import (
	"database/sql"
	"fmt"
	"net/url"
)

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	SubCategory string  `json:"subCategory"`
	Brand       string  `json:"brand"`
}

// func getTotalQueryResult(keyword string) int {
// 	var cfg = mysql.Config{
// 		User:   "myuser",
// 		Passwd: "mypassword",
// 		Net:    "tcp",
// 		Addr:   "localhost:3306",
// 		DBName: "myapp",
// 	}

// 	var err error
// 	s.db, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := s.db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	log.Print("Db is Connected")
// }

func Search(query url.Values, db *sql.DB) ([]Product, error) {
	keyword := "%" + query["q"][0] + "%"
	fmt.Println(query)

	q := `	SELECT
							p.name,
							p.description,
							p.price,
							b.name,
							c.name,
							sc.name
					FROM
							product p
					INNER JOIN category c ON
							c.category_id = p.category_id
					INNER JOIN brand b ON
							b.brand_id = p.brand_id
					INNER JOIN sub_category sc ON
							sc.sub_category_id = p.sub_category_id
					WHERE p.name LIKE ?`

	rows, err := db.Query(q, keyword)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Name, &product.Description, &product.Price, &product.Brand, &product.Category, &product.SubCategory); err != nil {
			return nil, fmt.Errorf("%s", err)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return products, nil
}
