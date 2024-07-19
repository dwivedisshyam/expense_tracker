package model

import "gofr.dev/pkg/gofr/http"

type Category struct {
	ID     string `bson:"id" json:"id"`
	UserID string `bson:"user_id" json:"user_id"`
	Name   string `bson:"name" json:"name"`
}

func (c Category) Validate() error {
	var params []string

	if c.Name == "" {
		params = append(params, "name")
	}

	if c.UserID == "" {
		params = append(params, "user_id")
	}

	if len(params) > 0 {
		return http.ErrorMissingParam{Params: params}
	}

	return nil
}

type CategoryFilter struct {
	ID     string
	UserID string
}
