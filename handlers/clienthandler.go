package handlers

import (
	"apigo/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type customerHandler struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustumoerHandler(usecase usecase.CustomerUsecase) customerHandler {
	return customerHandler{
		customerUsecase: usecase,
	}
}

func (c *customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customers, err := c.customerUsecase.GetCustomers()
	if err != nil {
		c.json(w, http.StatusInternalServerError, err)
	}
	c.json(w, http.StatusOK, customers)
}
