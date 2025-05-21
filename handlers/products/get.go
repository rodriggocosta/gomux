package products

import (
	"apigo/usecase/products"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	productCresteUsecase products.ProductUsecase
}

func NewProductHandler(usecase products.ProductUsecase) ProductHandler {
	return ProductHandler{
		productCresteUsecase: usecase,
	}
}

func (pr *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product, err := pr.productCresteUsecase.GetProduct()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(product)
}
