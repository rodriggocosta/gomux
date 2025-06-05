package products

import (
	"database/sql"
	"fmt"
)

type DeleteProdcutRepository struct {
	connection *sql.DB
}

func NewDeleteProductRepository(connection *sql.DB) DeleteProdcutRepository {
	return DeleteProdcutRepository{
		connection: connection,
	}
}

func (pr *DeleteProdcutRepository) Delete(product_id int) error {
	query := "DELETE FROM products WHERE product_id = $1"

	res, err := pr.connection.Exec(query, product_id)
	if err != nil {
		return fmt.Errorf("Erro ao excluir produto: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao excluir: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
