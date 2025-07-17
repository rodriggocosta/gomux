package address

import (
	"apigo/entity"
	"database/sql"
	"fmt"
)

type AddressPostReropsitory struct {
	connection *sql.DB
}

func NewAddressCreate(connection *sql.DB) AddressPostReropsitory {
	return AddressPostReropsitory{
		connection: connection,
	}
}

func (ad *AddressPostReropsitory) Create(address entity.Address) (int, error) {
	var id_address int

	query, err := ad.connection.Prepare("INSERT INTO address(street, number, complement, neighborhood, city, state, zip_code, country, address_type, is_default) VALUES($1, $2, $3, $4, $5,$6, $7, $8, $9, $10 )RETURNING id_address")

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	err = query.QueryRow(address.Street, address.Number, address.Complement, address.Neighborhood, address.City, address.State, address.Zip_code, address.Country, address.Address_type, address.Is_default).Scan(&id_address)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	query.Close()

	return id_address, nil
}
