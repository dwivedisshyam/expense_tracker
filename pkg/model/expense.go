package model

type Expense struct {
	ID         int64   `json:"id"`
	UserID     int64   `json:"-"`
	CategoryID int64   `json:"category_id"`
	Title      string  `json:"title"`
	Amount     float64 `json:"amount"`
	DueDate    string  `json:"due_date"`
	Paid       bool    `json:"paid"`
}

type ExpFilter struct {
	UserID    int64
	StartDate string
	EndDate   string
}
