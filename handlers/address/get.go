package address

import (
	"apigo/usecase/address"
	"encoding/json"
	"net/http"
)

type GetAddressHandler struct {
	getAddressUsecase address.GetAddressUsecase
}

func NewGetAddresshandler(usecase address.GetAddressUsecase) GetAddressHandler {
	return GetAddressHandler{
		getAddressUsecase: usecase,
	}
}

func (ad *GetAddressHandler) GetAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	address, err := ad.getAddressUsecase.GetAddress()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(address)
}
