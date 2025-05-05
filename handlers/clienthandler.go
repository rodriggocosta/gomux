package handlers

import (
	"apigo/entity"
	"apigo/usecase"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	w.Header().Set("Content-Type", "application/json")
	var customer entity.Customers
	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	insertCustomer, err := c.customerUsecase.PostCustomer(customer)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(insertCustomer)
}

func (c *CustomerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idString := r.URL.Query().Get("customer_id")
	customID, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "ID do cliente e obrigatorio", http.StatusBadRequest)
		return
	}

	customer, err := c.customerUsecase.GetCustomerById(customID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Cliente nao encontrado", http.StatusNotFound)
			return
		}

		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		fmt.Println("Erro ao codificar JSON", err)
		http.Error(w, "Erro ao retornar dados do cliente", http.StatusInternalServerError)
	}

}
