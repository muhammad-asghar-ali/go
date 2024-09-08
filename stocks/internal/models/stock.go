package models

import "time"

type (
	Stock struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		Price     string    `json:"price"`
		Company   string    `json:"company"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
