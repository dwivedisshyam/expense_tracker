package store

import (
	"database/sql"
	goErr "errors"

	"github.com/dwivedisshyam/expense_tracker/db"
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
)

type incomeStore struct {
	db *db.DB
}

func NewIncome(db *db.DB) Income {
	return &incomeStore{db}
}

func (us *incomeStore) Create(i *model.Income) (*model.Income, error) {
	q := `INSERT INTO incomes (user_id, title, amount, date) VALUES ($1,$2,$3,$4)`

	result, err := us.db.Exec(q, i.UserID, i.Title, i.Amount, i.Date)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	i.ID = id

	return i, nil
}

func (us *incomeStore) Update(i *model.Income) (*model.Income, error) {
	q := `UPDATE incomes set title=$1, amount=$2, date=$3 WHERE id=$4 AND user_id=$5`

	_, err := us.db.Exec(q, i.Title, i.Amount, i.Date, i.ID, i.UserID)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return i, nil
}

func (us *incomeStore) Get(i *model.Income) (*model.Income, error) {
	q := `SELECT title,amount,date FROM incomes WHERE id=$1 AND user_id=$2`

	err := us.db.QueryRow(q, i.ID, i.UserID).Scan(&i.Title, &i.Amount, &i.Date)
	if err != nil {
		if goErr.Is(err, sql.ErrNoRows) {
			return nil, errors.NotFound("user not found")
		}

		return nil, errors.Unexpected(err.Error())
	}

	return i, nil
}

func (us *incomeStore) Delete(i *model.Income) error {
	q := `DELETE FROM incomes WHERE id=$1 AND user_id=$2`

	_, err := us.db.Exec(q, i.ID, i.UserID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}
