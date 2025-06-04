package products

import (
	"apigo/entity"
	"database/sql"

	"github.com/shopspring/decimal"
)

type ProductCreateRepository struct {
	connection *sql.DB
}

func NewProductsCreateRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

// refatorar essa parte amanha
func (pr *ProductCreateRepository) Create(products entity.Products) (int, error) {
	priceDec, err := decimal.NewFromString(products.Price)
	if err != nil {
		return 0, err
	}
	query, err := pr.connection.Prepare("INSERT INTO products(name, price, code, validity, stock, entrace) VALUES ($1, $2, $3, $4, $5, $6) RETURNING product_id")
	if err != nil {
		return 0, err
	}
	defer query.Close()

	var productID int
	err = query.QueryRow(
		products.Name,
		priceDec.String(),
		products.Code,
		products.Validity,
		products.Stock,
		products.Entrace,
	).Scan(&productID)

	if err != nil {
		return 0, err
	}
	return productID, nil
}
