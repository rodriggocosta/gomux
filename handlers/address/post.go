package address

import (
	"apigo/entity"
	"apigo/usecase/address"
	"encoding/json"
	"net/http"
)

type AddressHandler struct {
	addressusecase address.AddressUsecase
}

func NewAddressHandler(usecase address.AddressUsecase) AddressHandler {
	return AddressHandler{
		addressusecase: usecase,
	}
}

func (ad *AddressHandler) Create(w http.ResponseWriter, res *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var address entity.Address
	err := json.NewDecoder(res.Body).Decode(&address)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	insertAddress, err := ad.addressusecase.Create(address)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(insertAddress)
}
