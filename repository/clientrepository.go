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
			&customerObjt.Customer_id,
			&customerObjt.Name,
			&customerObjt.Email,
			&customerObjt.Phone,
			&customerObjt.Cpf,
			&customerObjt.Cnpj,
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

	query, err := cr.connection.Prepare("INSERT INTO customers(name_customer, email, phone, cpf, cnpj) VALUES($1, $2, $3, $4, $5) RETURNING customer_id")

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	// comentario temporario: TALVEZ EU TENHA QUE COLOCAR OS CAMPOS CREATEDAT E UPDATEAT
	err = query.QueryRow(customer.Name, customer.Email, customer.Phone, customer.Cpf, customer.Cnpj).Scan(&customer_id)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	query.Close()

	return customer_id, nil
}

func (cr *CustomerRepository) GetCustomerById(customer_id int) (*entity.Customers, error) {
	query, err := cr.connection.Prepare("SELECT * FROM customers WHERE customer_id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()
	var customer entity.Customers

	err = query.QueryRow(customer_id).Scan(
		&customer.Customer_id,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.Cpf,
		&customer.Cnpj,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &customer, nil

}

func (cr *CustomerRepository) DeleteById(customer_id int) error {
	// essa tambem e outra forma, e funciona de boas
	query := "DELETE FROM customers WHERE customer_id = $1"

	res, err := cr.connection.Exec(query, customer_id)
	if err != nil {
		return fmt.Errorf("Erro ao deletar cliente: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao verifcar delecao: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (cr *CustomerRepository) Update(customer_id int, customer *entity.Customers) error {
	//	query, err := cr.connection.Prepare("UPADATE customers SET name = $1, email = $2, phone = $3, customer_id = %4")
	query := "UPDATE customers SET name_customer = $1, email = $2, phone = $3, cpf = $4, cnpj =$5 WHERE customer_id = $6"

	res, err := cr.connection.Exec(query, customer.Name, customer.Email, customer.Phone, customer.Cpf, customer.Cnpj, customer_id)

	if err != nil {
		return fmt.Errorf("Erro ao atualizar dados: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Erro ao tentar atualizar: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
