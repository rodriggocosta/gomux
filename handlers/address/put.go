package address

import (
	"apigo/entity"
	"apigo/usecase/address"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type PutAddressHandler struct {
	addressUsecase address.PutAddressUsecase
}

func NewPutAddressHandler(usecase address.PutAddressUsecase) PutAddressHandler {
	return PutAddressHandler{
		addressUsecase: usecase,
	}
}

func (ad *PutAddressHandler) PutAddress(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	conversaoId := req.URL.Query().Get("id_address")
	addressID, err := strconv.Atoi(conversaoId)

	if err != nil {
		http.Error(w, "ID do endereco obrigatorio", http.StatusInternalServerError)
		return
	}

	var address entity.Address

	if err := json.NewDecoder(req.Body).Decode(&address); err != nil {
		http.Error(w, "Formato Invalido", http.StatusBadRequest)
		return
	}

	err = ad.addressUsecase.PutAddress(addressID, &address)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Enderenco nao encontrado", http.StatusNotFound)
			return
		}
		log.Printf("Error ao atualizar os dados: %v", err)
		http.Error(w, "Erro No servidor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"Mensagem": "Atualizado com sucesso!"})

}
