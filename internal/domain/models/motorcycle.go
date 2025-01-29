package models

import "time"

type Motorcycle struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	Price     float64   `json:"price"`
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
}
