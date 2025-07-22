package address

import (
	"apigo/entity"
	"database/sql"
	"fmt"
)

type GetAdrressRepository struct {
	connection *sql.DB
}

func NewGetAddress(connection *sql.DB) GetAdrressRepository {
	return GetAdrressRepository{
		connection: connection,
	}
}

func (ad *GetAdrressRepository) GetAddress() ([]entity.Address, error) {
	query := "SELECT * FROM address"
	rows, err := ad.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	addressList := []entity.Address{}
	var addressObj entity.Address

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&addressObj.ID_Address,
			&addressObj.Street,
			&addressObj.Number,
			&addressObj.Complement,
			&addressObj.Neighborhood,
			&addressObj.City,
			&addressObj.State,
			&addressObj.Zip_code,
			&addressObj.Country,
			&addressObj.Address_type,
			&addressObj.Is_default,
			&addressObj.CreatedAt,
			&addressObj.UpdateAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err

		}

		addressList = append(addressList, addressObj)

	}

	return addressList, nil

}
