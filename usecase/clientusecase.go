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
