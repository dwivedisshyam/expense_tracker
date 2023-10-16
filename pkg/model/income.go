package model

import "time"

type Income struct {
	ID     int64     `json:"id"`
	UserID int64     `json:"user_id"`
	Title  string    `json:"title"`
	Amount float64   `json:"amount"`
	Date   time.Time `json:"date"`
}
