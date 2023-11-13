package model

import "github.com/dwivedisshyam/go-lib/pkg/errors"

type User struct {
	ID       int64  `json:"id"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u User) Validate() error {
	if u.FName == "" {
		return errors.Validation("f_name missing")
	}

	if u.LName == "" {
		return errors.Validation("l_name missing")
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
	ID    int64
	Email string
}
