package entity

import "time"

type Customers struct {
	ID int `json:"costumer_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
