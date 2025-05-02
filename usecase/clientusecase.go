package usecase

import (
	"apigo/entity"
	"apigo/repository"
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
	customer.ID = customerId

	return customer, nil
}
