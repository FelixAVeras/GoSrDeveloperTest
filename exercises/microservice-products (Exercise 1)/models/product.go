package models

import "time"

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Description string    `json:"description"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}