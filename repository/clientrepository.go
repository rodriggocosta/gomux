package repository

import (
	"apigo/entity"
	"database/sql"
	"fmt"
)

type CustomerRepository struct {
	connection *sql.DB
}

func NewCustomerRepository(connection *sql.DB) CustomerRepository {
	return CustomerRepository{
		connection: connection,
	}
}

func (cr *CustomerRepository) GetCustomers() ([]entity.Customers, error) {
	query := "SELECT * FROM customers"
	rows, err := cr.connection.Query(query)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	clientList := []entity.Customers{}
	var customerObjt entity.Customers

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&customerObjt.ID,
			&customerObjt.Name,
			&customerObjt.Email,
			&customerObjt.Phone,
			&customerObjt.CreatedAt,
			&customerObjt.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		clientList = append(clientList, customerObjt)

	}

	return clientList, nil
}

func (cr *CustomerRepository) PostCustomer(customer entity.Customers) (int, error) {
	var customer_id int

	query, err := cr.connection.Prepare("INSERT INTO customers(name, email, phone) VALUES ($1, $2, $3) RETURNING customer_id")

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	// comentario temporario: TALVEZ EU TENHA QUE COLOCAR OS CAMPOS CREATEDAT E UPDATEAT
	err = query.QueryRow(customer.Name, customer.Email, customer.Phone).Scan(&customer_id)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	query.Close()

	return customer_id, nil
}
