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
			&productsObj.Product_id,
			&productsObj.Name,
			&productsObj.Price,
			&productsObj.Code,
			&productsObj.Validity,
			&productsObj.Stock,
			&productsObj.Entrace,
			&productsObj.Createat,
			&productsObj.Updatedat,
		)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		productsList = append(productsList, productsObj)
	}
	return productsList, nil
}
