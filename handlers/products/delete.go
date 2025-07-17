package products

import (
	"apigo/usecase/products"
	"database/sql"
	"net/http"
	"strconv"
)

type DeleteProductHandler struct {
	usecasehandler products.DeleteProductUsecase
}

func NewDeleteProductHandler(usecase products.DeleteProductUsecase) DeleteProductHandler {
	return DeleteProductHandler{
		usecasehandler: usecase,
	}
}

func (pr *DeleteProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	convert := r.URL.Query().Get("product_id")
	productID, err := strconv.Atoi(convert)

	if err != nil {
		http.Error(w, "ID do produto obrigatorio", http.StatusBadRequest)
		return
	}

	err = pr.usecasehandler.Delete(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Produto nao encontrado na base de dados", http.StatusNotFound)
			return
		}

		http.Error(w, "Erro no Servidor", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
