package products

import (
	"apigo/entity"
	"database/sql"
	"fmt"
)

type ByIdProductRepository struct {
	connection *sql.DB
}

func NewByIdProductRepository(connection *sql.DB) ByIdProductRepository {
	return ByIdProductRepository{
		connection: connection,
	}
}

func (pr *ByIdProductRepository) ByIdProduct(product_id int) (*entity.Products, error) {
	query, err := pr.connection.Prepare("SELECT * FROM products WHERE product_id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	var product entity.Products

	err = query.QueryRow(product_id).Scan(
		&product.Product_id,
		&product.Name,
		&product.Price,
		&product.Code,
		&product.Validity,
		&product.Stock,
		&product.Entrace,
		&product.Createat,
		&product.Updatedat,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return &product, nil
}
