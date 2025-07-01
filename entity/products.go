package entity

import (
	"time"
)

type Products struct {
	Product_id int       `Json:"Product_id"`
	Name       string    `json:"name_product"`
	Price      string    `json:"price"`
	Code       int       `json:"code"`
	Validity   time.Time `json:"validity"`
	Stock      int       `json:"stock"`
	Entrace    time.Time `json:"entrace"`
	Createat   time.Time `json:"createat"`
	Updatedat  time.Time `json:"upatedat"`
}
