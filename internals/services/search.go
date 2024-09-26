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

func Search(query url.Values, db *sql.DB) ([]Product, error) {
	keyword := query["q"][0]

	q := `	select
							p.name,
							p.description,
							p.price,
							b.name,
							c.name,
							sc.name
					from
							product p
					inner join category c on
							c.category_id = p.category_id
					inner join brand b on
							b.brand_id = p.brand_id
					inner join sub_category sc on
							sc.sub_category_id = p.sub_category_id
					WHERE p.name = ?`

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
