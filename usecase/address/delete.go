package address

import "apigo/repository/address"

type DeleteAddressUsecase struct {
	repostitory address.DeleteAddressRepository
}

func NewDeleteAddress(repo address.DeleteAddressRepository) DeleteAddressUsecase {
	return DeleteAddressUsecase{
		repostitory: repo,
	}
}

func (ad *DeleteAddressUsecase) Delete(id_address int) error {
	return ad.repostitory.Delete(id_address)
}
