package entity

import "time"

type Address struct {
	ID_Address   int       `json:"id_address"`
	Street       string    `json:"street"`
	Number       int       `json:"number"`
	Complement   string    `json:"complement"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Zip_code     string    `json:"zip_code"`
	Address_type string    `json:"address_type"`
	Country      string    `json:"country"`
	Is_default   bool      `json:"is_default"`
	CreatedAt    time.Time `json:"created_at"`
	UpdateAt     time.Time `json:"updated_at"`
}
