package products

import (
	"apigo/entity"
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
)

type PutProductRepository struct {
	connection *sql.DB
}

func NewPutProductRepository(connection *sql.DB) PutProductRepository {
	return PutProductRepository{
		connection: connection,
	}
}

func (pr *PutProductRepository) PutProduct(product_id int, product *entity.Products) error {

	priceDec, err := decimal.NewFromString(product.Price)
	if err != nil {
		return fmt.Errorf("Erro ao converter o preco: %w", err)
	}

	query := "UPDATE products SET name = $1, price = $2, code = $3, validity = $4, stock = $5, entrace = $6  WHERE product_id = $7"

	resp, err := pr.connection.Exec(
		query,
		product.Name,
		priceDec,
		product.Code,
		product.Validity,
		product.Stock,
		product.Entrace,
		product_id,
	)

	if err != nil {
		return fmt.Errorf("Erro ao tentar atualizar: %w", err)
	}

	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao verificar atualizacao: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
