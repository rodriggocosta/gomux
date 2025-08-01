package address

import (
	"database/sql"
	"fmt"
)

type DeleteAddressRepository struct {
	connection *sql.DB
}

func NewDeleteAddressRespository(connection *sql.DB) DeleteAddressRepository {
	return DeleteAddressRepository{
		connection: connection,
	}
}

func (ad *DeleteAddressRepository) Delete(id_address int) error {
	query := "DELETE FROM address WHERE id_address = $1"
	res, err := ad.connection.Exec(query, id_address)
	if err != nil {
		return fmt.Errorf("Erro ao delete endereco: %w", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao verificar delecao: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
