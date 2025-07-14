package products

import (
	"apigo/usecase/products"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ByIdProductHandler struct {
	produtcByIdUsecase products.ProductByIdUsecase
}

func NewProductByIdHandler(usecase products.ProductByIdUsecase) ByIdProductHandler {
	return ByIdProductHandler{
		produtcByIdUsecase: usecase,
	}
}

func (pr *ByIdProductHandler) GetByIdProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idString := r.URL.Query().Get("product_id")
	productID, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "ID do cliente e obrigatorio", http.StatusBadRequest)
		return
	}

	product, err := pr.produtcByIdUsecase.GetByIdProduct(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Cliente nao encontrado", http.StatusNotFound)
			return
		}

		http.Error(w, "Erro no servidor", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		fmt.Println("Erro ao codificar o JSON", err)
		http.Error(w, "Erro ao retornar dados do cliente", http.StatusInternalServerError)
	}

}
