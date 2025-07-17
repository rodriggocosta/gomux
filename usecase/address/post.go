package address

import (
	"apigo/entity"
	"apigo/repository/address"
)

type AddressUsecase struct {
	respository address.AddressPostReropsitory
}

func NewAddress(repo address.AddressPostReropsitory) AddressUsecase {
	return AddressUsecase{
		respository: repo,
	}
}

func (ad *AddressUsecase) Create(address entity.Address) (entity.Address, error) {
	addressID, err := ad.respository.Create(address)
	if err != nil {
		return entity.Address{}, err
	}
	address.ID_Address = addressID
	return address, nil
}
