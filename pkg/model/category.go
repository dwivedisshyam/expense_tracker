package model

type Category struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

type CatFilter struct {
	UserID int64
}
