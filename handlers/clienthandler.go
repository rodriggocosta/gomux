package handlers

import (
	"apigo/entity"
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

func (c *CustomerHandler) PostCustomer(w http.ResponseWriter, r *http.Request) {
	var customer entity.Customers
	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	insertCustomer, err := c.customerUsecase.PostCustomer(customer)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(insertCustomer)

	w.WriteHeader(http.StatusCreated)

}
