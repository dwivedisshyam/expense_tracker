package model

import "gofr.dev/pkg/gofr/http"

type Expense struct {
	ID         string  `bson:"id" json:"id"`
	UserID     string  `bson:"user_id" json:"user_id"`
	CategoryID string  `bson:"category_id" json:"category_id"`
	Title      string  `bson:"title" json:"title"`
	Amount     float64 `bson:"amount" json:"amount"`
	DueDate    string  `bson:"due_date" json:"due_date"`
	Paid       bool    `bson:"paid" json:"paid"`
}

func (e Expense) Validate() error {
	var params []string

	if e.Title == "" {
		params = append(params, "title")
	}

	if e.Amount < 0 {
		params = append(params, "amount")
	}

	if e.CategoryID == "" {
		params = append(params, "category_id")
	}

	if e.UserID == "" {
		params = append(params, "user_id")
	}

	if len(params) > 0 {
		return http.ErrorMissingParam{Params: params}
	}

	return nil
}

type ExpenseFilter struct {
	ID        string
	UserID    string
	StartDate string
	EndDate   string
}
