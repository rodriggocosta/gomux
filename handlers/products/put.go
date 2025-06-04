package products

import (
	"apigo/entity"
	"apigo/usecase/products"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type PutProductHandler struct {
	productUsercase products.PutProductUsecase
}

func NewPutProudcHandler(usecase products.PutProductUsecase) PutProductHandler {
	return PutProductHandler{
		productUsercase: usecase,
	}
}

func (pr *PutProductHandler) PutProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	converterId := r.URL.Query().Get("product_id")
	productID, err := strconv.Atoi(converterId)

	if err != nil {
		http.Error(w, "Id do produto obrigatorio", http.StatusBadRequest)
		return
	}

	var product entity.Products

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Formato Invalido", http.StatusBadRequest)
		return
	}

	err = pr.productUsercase.PutProduct(productID, &product)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Produto nao encontrado", http.StatusNotFound)
			return
		}
		log.Printf("Erro ao atualizar produtos: %v", err)
		http.Error(w, "Error no servidor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Produto atualizado"})
}
