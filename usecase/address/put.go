package address

import (
	"apigo/entity"
	"apigo/repository/address"
)

type PutAddressUsecase struct {
	repository address.UpdateAddressRepository
}

func NewPutAddress(repo address.UpdateAddressRepository) PutAddressUsecase {
	return PutAddressUsecase{
		repository: repo,
	}
}

func (ad *PutAddressUsecase) PutAddress(id_address int, address *entity.Address) error {
	return ad.repository.PutAddress(id_address, address)
}
