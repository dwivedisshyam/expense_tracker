package sqlite

import (
	"database/sql"
	goErr "errors"
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type expenseStore struct {
	idGen func(time.Time) string
}

func NewExpense(idGen func(time.Time) string) *expenseStore {
	return &expenseStore{idGen: idGen}
}

func (se expenseStore) Index(ctx *gofr.Context, f *model.ExpenseFilter) ([]model.Expense, error) {
	var exps []model.Expense

	c, args := clause(f)

	q := `SELECT id,title,amount,due_date,category_id,user_id,is_paid FROM expenses WHERE ` + c

	rows, err := ctx.SQL.Query(q, args...)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	for rows.Next() {
		var e model.Expense

		err = rows.Scan(&e.ID, &e.Title, &e.Amount, &e.DueDate, &e.CategoryID, &e.UserID, &e.Paid)
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
func (es expenseStore) Create(ctx *gofr.Context, exp *model.Expense) (*model.Expense, error) {
	q := `INSERT INTO expenses (id, user_id, title, amount,category_id,due_date,is_paid) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	exp.ID = es.idGen(time.Now())
	_, err := ctx.SQL.Exec(q, exp.ID, exp.UserID, exp.Title, exp.Amount, exp.CategoryID, exp.DueDate, exp.Paid)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return exp, nil
}
func (se expenseStore) Update(ctx *gofr.Context, exp *model.Expense) error {
	q := `UPDATE expenses set title=$1, amount=$2,due_date=$3, category_id=$4, is_paid=$5 WHERE id=$6 AND user_id=$7`

	_, err := ctx.SQL.Exec(q, exp.Title, exp.Amount, exp.DueDate, exp.CategoryID, exp.Paid, exp.ID, exp.UserID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}
func (se expenseStore) Get(ctx *gofr.Context, f *model.ExpenseFilter) (*model.Expense, error) {
	q := `SELECT id,user_id,title,amount,due_date,category_id,is_paid FROM expenses WHERE id=$1 AND user_id=$2`

	var e model.Expense

	err := ctx.SQL.QueryRow(q, f.ID, f.UserID).Scan(&e.ID, &e.UserID, &e.Title, &e.Amount, &e.DueDate, &e.CategoryID, &e.Paid)
	if err != nil {
		if goErr.Is(err, sql.ErrNoRows) {
			return nil, errors.NotFound("user not found")
		}

		return nil, errors.Unexpected(err.Error())
	}

	return &e, nil
}
func (se expenseStore) Delete(ctx *gofr.Context, f *model.ExpenseFilter) error {
	q := `DELETE FROM expenses WHERE id=$1 AND user_id=$2`

	_, err := ctx.SQL.Exec(q, f.ID, f.UserID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func clause(f *model.ExpenseFilter) (c string, args []interface{}) {
	c = `user_id=$1 `

	args = append(args, f.UserID)

	if f.StartDate != "" && f.EndDate != "" {
		c += `AND (due_date BETWEEN $2 AND $3) `

		args = append(args, f.StartDate, f.EndDate+" 23:59:59")
	}

	return
}
