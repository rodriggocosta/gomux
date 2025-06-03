package products

import (
	"apigo/entity"
	"database/sql"
	"fmt"
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
	query := "UPDATE products SET name = $1, price = $2, code = $3, validity = $4, stock = $5, entrace = $6"

	resp, err := pr.connection.Exec(query, product.Name, product.Code, product.Price, product.Validity, product.Stock, product.Entrace)

	if err != nil {
		return fmt.Errorf("Erro ao atauliazar dados: %w", err)
	}

	rowsAffcted, err := resp.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao tentar atualizar: %w", err)
	}

	if rowsAffcted == 0 {
		return sql.ErrNoRows
	}

	return nil
}
