package address

import (
	"apigo/entity"
	"apigo/repository/address"
)

type GetAddressUsecase struct {
	repository address.GetAdrressRepository
}

func NewAddressUsecase(repo address.GetAdrressRepository) GetAddressUsecase {
	return GetAddressUsecase{
		repository: repo,
	}
}

func (ad *GetAddressUsecase) GetAddress() ([]entity.Address, error) {
	return ad.repository.GetAddress()
}
