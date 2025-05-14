package entity

import (
	"time"
)

type Products struct {
	ProductID     int       `json:"produto_id"`
	Name          string    `json:"name"`
	Code          int       `json:"code"`
	Price         float64   `json:"price"`
	Stock         int       `json:"stock"`
	Category      string    `json:"category"`
	Date_Validity time.Time `json:"date_validity"`
	Brand         string    `json:"brand"`
	Sector        string    `json:"sector"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"upated_at"`
	CustomerID    int       `json:"customer_id"`
}
