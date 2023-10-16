package model

type User struct {
	ID       int64  `json:"id"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserFilter struct {
	ID    int64
	Email string
}
