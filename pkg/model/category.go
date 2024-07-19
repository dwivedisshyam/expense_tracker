package model

type Category struct {
	ID     string `bson:"id" json:"id"`
	UserID string `bson:"user_id" json:"user_id"`
	Name   string `bson:"name" json:"name"`
}

type CategoryFilter struct {
	ID     string
	UserID string
}
