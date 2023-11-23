package store

import (
	"database/sql"
	goErr "errors"

	"github.com/dwivedisshyam/expense_tracker/db"
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
)

type expStore struct {
	db *db.DB
}

func NewExpense(db *db.DB) Expense {
	return &expStore{db}
}

func (us *expStore) Index(f *model.ExpFilter) ([]model.Expense, error) {
	var exps []model.Expense

	c, args := clause(f)

	q := `SELECT id,title,amount,due_date,category_id,is_paid FROM expenses WHERE ` + c

	rows, err := us.db.Query(q, args...)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	for rows.Next() {
		var e model.Expense

		err = rows.Scan(&e.ID, &e.Title, &e.Amount, &e.DueDate, &e.CategoryID, &e.Paid)
		if err != nil {
			return nil, errors.Unexpected(err.Error())
		}

		exps = append(exps, e)
	}

	if rows.Err() != nil {
		return nil, errors.Unexpected("DB Error")
	}

	return exps, nil
}

func (us *expStore) Create(e *model.Expense) (*model.Expense, error) {
	q := `INSERT INTO expenses (user_id, title, amount,category_id,due_date,is_paid) VALUES ($1,$2,$3,$4,$5,$6)`

	_, err := us.db.Exec(q, e.UserID, e.Title, e.Amount, e.CategoryID, e.DueDate, e.Paid)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return e, nil
}

func (us *expStore) Update(e *model.Expense) (*model.Expense, error) {
	q := `UPDATE expenses set title=$1, amount=$2,due_date=$3, category_id=$4, is_paid=$5 WHERE id=$6 AND user_id=$7`

	_, err := us.db.Exec(q, e.Title, e.Amount, e.DueDate, e.CategoryID, e.Paid, e.ID, e.UserID)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return e, nil
}

func (us *expStore) Get(e *model.Expense) (*model.Expense, error) {
	q := `SELECT title,amount,due_date,category_id,is_paid FROM expenses WHERE id=$1 AND user_id=$2`

	err := us.db.QueryRow(q, e.ID, e.UserID).Scan(&e.Title, &e.Amount, &e.DueDate, &e.CategoryID, &e.Paid)
	if err != nil {
		if goErr.Is(err, sql.ErrNoRows) {
			return nil, errors.NotFound("user not found")
		}

		return nil, errors.Unexpected(err.Error())
	}

	return e, nil
}

func (us *expStore) Delete(e *model.Expense) error {
	q := `DELETE FROM expenses WHERE id=$1 AND user_id=$2`

	_, err := us.db.Exec(q, e.ID, e.UserID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func clause(f *model.ExpFilter) (c string, args []interface{}) {
	c = `user_id=$1 `

	args = append(args, f.UserID)

	if f.StartDate != "" && f.EndDate != "" {
		c += `AND (due_date BETWEEN $2 AND $3) `

		args = append(args, f.StartDate, f.EndDate+" 23:59:59")
	}

	return
}
