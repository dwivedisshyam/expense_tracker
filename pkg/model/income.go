package model

import "time"

type Income struct {
	ID     string    `json:"id"`
	UserID string    `json:"user_id"`
	Title  string    `json:"title"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}

type IncomeFilter struct {
}
