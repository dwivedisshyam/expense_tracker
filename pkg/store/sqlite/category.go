package sqlite

import (
	"database/sql"
	goErr "errors"
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type categoryStore struct {
	idGen func(time.Time) string
}

func NewCategory(idGen func(time.Time) string) *categoryStore {
	return &categoryStore{idGen: idGen}
}

func (cs categoryStore) Index(ctx *gofr.Context, f *model.CategoryFilter) ([]model.Category, error) {
	q := `SELECT id,name,user_id FROM categories WHERE user_id=$1`

	rows, err := ctx.SQL.Query(q, f.UserID)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	defer rows.Close()

	var cats []model.Category

	for rows.Next() {
		var c model.Category

		err = rows.Scan(&c.ID, &c.Name, &c.UserID)
		if err != nil {
			return nil, errors.Unexpected(err.Error())
		}

		cats = append(cats, c)
	}

	if rows.Err() != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return cats, nil
}

func (sc categoryStore) Create(ctx *gofr.Context, cat *model.Category) (*model.Category, error) {
	q := `INSERT INTO categories (id, name, user_id) VALUES ($1,$2,$3)`

	cat.ID = sc.idGen(time.Now())

	_, err := ctx.SQL.Exec(q, cat.ID, cat.Name, cat.UserID)
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}

	return cat, nil
}

func (sc categoryStore) Update(ctx *gofr.Context, cat *model.Category) error {
	q := `UPDATE categories set name=$1 WHERE id=$2 AND user_id=$3`

	_, err := ctx.SQL.Exec(q, cat.Name, cat.ID, cat.UserID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}

func (sc categoryStore) Get(ctx *gofr.Context, f *model.CategoryFilter) (*model.Category, error) {
	q := `SELECT id,user_id,name FROM categories WHERE id=$1 AND user_id=$2`

	var cat model.Category

	err := ctx.SQL.QueryRow(q, f.ID, f.UserID).Scan(&cat.ID, &cat.UserID, &cat.Name)
	if err != nil {
		if goErr.Is(err, sql.ErrNoRows) {
			return nil, errors.NotFound("user not found")
		}

		return nil, errors.Unexpected(err.Error())
	}

	return &cat, nil
}

func (sc categoryStore) Delete(ctx *gofr.Context, f *model.CategoryFilter) error {
	q := `DELETE FROM categories WHERE id=$1 AND user_id=$2`

	_, err := ctx.SQL.Exec(q, f.ID, f.UserID)
	if err != nil {
		return errors.Unexpected(err.Error())
	}

	return nil
}
