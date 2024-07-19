package model

import "gofr.dev/pkg/gofr/http"

type User struct {
	ID        string `bson:"id" json:"id"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password,omitempty"`
	// ProjectID string `bson:"project_id" json:"project_id,omitempty"`
}

func (u User) Validate() error {
	var params []string
	if u.FirstName == "" {
		params = append(params, "first_name")
	}

	if u.LastName == "" {
		params = append(params, "last_name")
	}

	if u.Email == "" {
		params = append(params, "email")
	}

	if u.Password == "" {
		params = append(params, "password")
	}

	if len(params) > 0 {
		return http.ErrorMissingParam{Params: params}
	}

	return nil
}

type UserFilter struct {
	ID    string
	Email string
}
