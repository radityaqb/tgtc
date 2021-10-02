package service

import (
	"errors"

	"github.com/radityaqb/tgtc/backend/database"
	"github.com/radityaqb/tgtc/backend/dictionary"
)

func GetProduct(paramID int) (*dictionary.Product, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT product_id, product_name, product_price, product_image, shop_name
	FROM products
	WHERE product_id = $1
	`
	// actual query process
	row := db.QueryRow(query, paramID)

	// read query result, and assign to variable(s)
	var result dictionary.Product
	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.ProductPrice,
		&result.ShopName,
		&result.ImageURL,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetProducts() ([]dictionary.Product, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT product_id, product_name, product_price, product_image, shop_name
	FROM products
	`
	// actual query process
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dictionary.Product
	for rows.Next() {
		var data dictionary.Product
		rows.Scan(
			&data.ID,
			&data.Name,
			&data.ProductPrice,
			&data.ImageURL,
			&data.ShopName,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func CreateProduct(data dictionary.Product) (*dictionary.Product, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	INSERT INTO products (product_name, product_price, product_image, shop_name) VALUES
		($1, $2, $3, $4)
	`
	// actual query process
	result, err := db.Exec(query,
		data.Name,
		data.ProductPrice,
		data.ImageURL,
		data.ShopName,
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row created")
	}

	return &data, nil
}

func UpdateProduct(data dictionary.Product) (*dictionary.Product, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	UPDATE 
		products
	SET 
		product_name = $2,
		product_price = $3,
		product_image = $4,
		shop_name = $5
	WHERE
		product_id = $1
	`
	// actual query process
	result, err := db.Exec(query,
		data.ID,
		data.Name,
		data.ProductPrice,
		data.ImageURL,
		data.ShopName,
	)
	if err != nil {
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row updated")
	}

	return &data, nil
}
