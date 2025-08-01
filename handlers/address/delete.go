package address

import (
	"apigo/usecase/address"
	"database/sql"
	"net/http"
	"strconv"
)

type DeleteAddressHandler struct {
	addressUsecase address.DeleteAddressUsecase
}

func NewDeleteAddressHandler(usecase address.DeleteAddressUsecase) DeleteAddressHandler {
	return DeleteAddressHandler{
		addressUsecase: usecase,
	}
}

func (ad *DeleteAddressHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	converterID := r.URL.Query().Get("id_address")
	addressID, err := strconv.Atoi(converterID)

	if err != nil {
		http.Error(w, "ID do endereco  obrigatorios", http.StatusBadRequest)
		return
	}

	err = ad.addressUsecase.Delete(addressID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Endereco nao econtrado", http.StatusNotFound)
			return
		}

		http.Error(w, "Erro no Servidor", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusNoContent)
}
