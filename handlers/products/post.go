package products

import (
	"apigo/entity"
	"apigo/usecase/products"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductCreateHandler struct {
	productUsecase products.ProductCreateUsecase
}

func NewProductsCreateHandlers(usecase products.ProductCreateUsecase) ProductCreateHandler {
	return ProductCreateHandler{
		productUsecase: usecase,
	}
}

func (pr *ProductCreateHandler) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var product entity.Products

	if product.Price == "" {
		fmt.Println("erro aqui")
	}

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	insertProducts, err := pr.productUsecase.Create(product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertProducts)

}
