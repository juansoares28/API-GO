package repository

import (
	"API/models"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

var id int

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT id, product_name, price  FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}
	var productsList []models.Product
	var productObj models.Product
	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)
		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}
	}
	productsList = append(productsList, productObj)
	rows.Close()
	return productsList, nil
}
func (pr *ProductRepository) CreateProduct(product models.Product) (int, error) {
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) returning id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}
