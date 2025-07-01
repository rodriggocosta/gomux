package entity

import "time"

type Customers struct {
	Customer_id int       `json:"customer_id"`
	Name        string    `json:"name_customer"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Cpf         string    `json:"cpf"`
	Cnpj        string    `json:"cnpj"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
