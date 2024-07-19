package model

type Expense struct {
	ID         string  `bson:"id" json:"id"`
	UserID     string  `bson:"user_id" json:"user_id"`
	CategoryID string  `bson:"category_id" json:"category_id"`
	Title      string  `bson:"title" json:"title"`
	Amount     float64 `bson:"amount" json:"amount"`
	Date       string  `bson:"due_date" json:"due_date"`
	Paid       bool    `bson:"paid" json:"paid"`
}

type ExpenseFilter struct {
	ID        string
	UserID    string
	StartDate string
	EndDate   string
}
