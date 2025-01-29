package models

import "time"

type Order struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	MotorcycleID int       `json:"motorcycle_id"`
	OrderDate    time.Time `json:"order_date"`
	TotalPrice   float64   `json:"total_price"`
	Quantity     int       `json:"quantity"`
}
