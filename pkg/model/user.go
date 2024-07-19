package model

import "github.com/dwivedisshyam/go-lib/pkg/errors"

type User struct {
	ID        string `bson:"id" json:"id"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password,omitempty"`
	// ProjectID string `bson:"project_id" json:"project_id,omitempty"`
}

func (u User) Validate() error {
	if u.FirstName == "" {
		return errors.Validation("first_name missing")
	}

	if u.LastName == "" {
		return errors.Validation("last_name missing")
	}

	if u.Email == "" {
		return errors.Validation("email missing")
	}

	if u.Password == "" {
		return errors.Validation("password missing")
	}

	return nil
}

type UserFilter struct {
	ID    string
	Email string
}
