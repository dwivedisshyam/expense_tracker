package store

import (
	"log"

	"github.com/dwivedisshyam/expense_tracker/db"
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
)

type userStore struct {
	db *db.DB
}

func NewUser(db *db.DB) User {
	return &userStore{db}
}

func (us *userStore) Create(user *model.User) (*model.User, error) {
	q := `INSERT INTO users (f_name, l_name, email, password) VALUES ($1,$2,$3,$4)`

	result, err := us.db.Exec(q, user.FName, user.LName, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}

func (us *userStore) Update(user *model.User) (*model.User, error) {
	q := `UPDATE users set f_name=$1, l_name=$2, email=$3, password=$4 WHERE id=$5`

	_, err := us.db.Exec(q, user.FName, user.LName, user.Email, user.Password, user.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (us *userStore) Get(f *model.UserFilter) (*model.User, error) {
	q := `SELECT id,f_name, l_name, email, password FROM users WHERE `

	var identifier any
	if f.Email != "" {
		identifier = f.Email
		q += `email=$1`
	} else {
		identifier = f.ID
		q += `id=$1`
	}

	user := new(model.User)
	err := us.db.QueryRow(q, identifier).Scan(&user.ID, &user.FName, &user.LName, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (us *userStore) Delete(user *model.User) error {
	q := `DELETE FROM users WHERE id=$1`

	_, err := us.db.Exec(q, user.ID)
	if err != nil {
		log.Println(err)
	}

	return err
}
