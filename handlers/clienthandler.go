package handlers

import (
	"apigo/usecase"
	"encoding/json"
	"net/http"
)

type CustomerHandler struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerHandler(usecase usecase.CustomerUsecase) CustomerHandler {
	return CustomerHandler{
		customerUsecase: usecase,
	}
}

func (c *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customers, err := c.customerUsecase.GetCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)

}
