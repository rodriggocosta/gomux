package usecase

import (
	"apigo/entity"
	"apigo/repository"
	"fmt"
)

type CustomerUsecase struct {
	repository repository.CustomerRepository
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return CustomerUsecase{
		repository: repo,
	}
}

func (cr *CustomerUsecase) GetCustomers() ([]entity.Customers, error) {
	return cr.repository.GetCustomers()
}

func (cr *CustomerUsecase) PostCustomer(customer entity.Customers) (entity.Customers, error) {
	customerId, err := cr.repository.PostCustomer(customer)

	if err != nil {
		return entity.Customers{}, err
	}
	customer.Customer_id = customerId

	return customer, nil
}

func (cr *CustomerUsecase) GetCustomerById(customer_id int) (*entity.Customers, error) {
	customer, err := cr.repository.GetCustomerById(customer_id)

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar cliente com ID %d: %w", customer_id, err)
	}

	return customer, nil
}

func (cr *CustomerUsecase) DeleteById(customer_id int) error {
	return cr.repository.DeleteById(customer_id)
}

func (cr *CustomerUsecase) Update(customer_id int, customer *entity.Customers) error {
	return cr.repository.Update(customer_id, customer)
}
