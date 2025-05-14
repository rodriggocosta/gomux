package products

import (
	"apigo/entity"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProduct() ([]entity.Products, error) {
	query := "SELECT * FROM products"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	productsList := []entity.Products{}
	var productsObj entity.Products

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&productsObj.ProductID,
			&productsObj.Name,
			&productsObj.Code,
			&productsObj.Price,
			&productsObj.Stock,
			&productsObj.Category,
			&productsObj.Date_Validity,
			&productsObj.Brand,
			&productsObj.Created_at,
			&productsObj.Updated_at,
			&productsObj.CustomerID,
			&productsObj.Sector,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		productsList = append(productsList, productsObj)
	}
	return productsList, nil
}
