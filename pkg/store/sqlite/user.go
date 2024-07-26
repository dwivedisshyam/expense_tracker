package sqlite

import (
	"database/sql"
	goErr "errors"
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type userStore struct {
	idGen func(time.Time) string
}

func NewUser(idGen func(time.Time) string) *userStore {
	return &userStore{idGen: idGen}
}

func (us userStore) Create(ctx *gofr.Context, user *model.User) (*model.User, error) {
	q := `INSERT INTO users (id, first_name, last_name, email, password) VALUES ($1,$2,$3,$4,$5)`

	user.ID = us.idGen(time.Now())

	_, err := ctx.SQL.Exec(q, user.ID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return user, nil
}

func (su userStore) Update(ctx *gofr.Context, user *model.User) error {
	q := `UPDATE users set first_name=$1, last_name=$2, email=$3, password=$4 WHERE id=$5`

	_, err := ctx.SQL.Exec(q, user.FirstName, user.LastName, user.Email, user.Password, user.ID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (su userStore) Get(ctx *gofr.Context, f *model.UserFilter) (*model.User, error) {
	q := `SELECT id, first_name, last_name, email, password FROM users WHERE `

	var identifier any
	if f.Email != "" {
		identifier = f.Email
		q += `email=$1`
	} else {
		identifier = f.ID
		q += `id=$1`
	}

	user := new(model.User)

	err := ctx.SQL.QueryRow(q, identifier).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if goErr.Is(err, sql.ErrNoRows) {
			return nil, errors.NotFound("user not found")
		}

		return nil, errors.Unexpected(err.Error())
	}

	return user, nil
}

func (su userStore) Delete(ctx *gofr.Context, f *model.UserFilter) error {
	q := `DELETE FROM users WHERE id=$1`

	if _, err := ctx.SQL.Exec(q, f.ID); err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}
