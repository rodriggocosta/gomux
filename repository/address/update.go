package address

import (
	"apigo/entity"
	"database/sql"
	"fmt"
)

type UpdateAddressRepository struct {
	connection *sql.DB
}

func NewUpdateAddressRepository(connection *sql.DB) UpdateAddressRepository {
	return UpdateAddressRepository{
		connection: connection,
	}
}

func (ad *UpdateAddressRepository) PutAddress(id_address int, address *entity.Address) {
	query := "UPDATE address SET street = $1, number = $2, complement = $3, neighborhood = $4, city = $5, state = $6, zip_code = $8, address_type = $9, country = $10, is_default = $11 WHERE id_address = $ 12"

	resp, err := ad.connection.Exec(
		query,
		address.Street,
		address.Number,
		address.Complement,
		address.Neighborhood,
		address.City,
		address.State,
		address.Zip_code,
		address.Address_type,
		address.Country,
		address.Is_default,
	)

	if err != nil {
		return fmt.Errorf("Erro ao atualizar: %w", err)
	}

	rowsAffected, err := resp.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao tentar atualizar: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil

}
