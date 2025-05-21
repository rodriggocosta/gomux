package products

import (
	"apigo/entity"
	"database/sql"
	"fmt"
)

type ProductCreateRepository struct {
	connection *sql.DB
}

func NewProductsCreateRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductCreateRepository) Create(products entity.Products) (int, error) {
	var product_id int
	query, err := pr.connection.Prepare("INSERT INTO products(name, code, price, stock, category, date_validity, brand, sector) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING customer_id")
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	// OBS: aqui eu tenho que ver dois a questao da tabela, por conta da fk de customers
	err = query.QueryRow(products.Name, products.Code, products.Price, products.Stock, products.Category, products.Date_Validity, products.Brand, products.Sector).Scan(&product_id)

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	defer query.Close()
	return product_id, nil
}
