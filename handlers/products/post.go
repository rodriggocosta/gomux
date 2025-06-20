package products

import (
	"apigo/entity"
	"apigo/usecase/products"
	"encoding/json"
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

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if product.Price == "" {
		http.Error(w, "Pricse ir required", http.StatusBadRequest)
		return
	}

	insertProducts, err := pr.productUsecase.Create(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertProducts)

}
