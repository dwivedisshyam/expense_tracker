package store

import (
	"github.com/dwivedisshyam/expense_tracker/db"
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
)

type categoryStore struct {
	db *db.DB
}

func NewCategory(db *db.DB) Category {
	return &categoryStore{db}
}

func (us *categoryStore) Index(f *model.CatFilter) ([]model.Category, error) {
	q := `SELECT id,name FROM categories WHERE user_id=$1`
	rows, err := us.db.Query(q, f.UserID)

	var cats []model.Category

	for rows.Next() {
		var c model.Category
		err = rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}

		cats = append(cats, c)
	}

	return cats, nil
}

func (us *categoryStore) Create(cat *model.Category) (*model.Category, error) {
	q := `INSERT INTO categories (name, user_id) VALUES ($1,$2)`

	result, err := us.db.Exec(q, cat.Name, cat.UserID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	cat.ID = id

	return cat, nil
}
func (us *categoryStore) Update(cat *model.Category) (*model.Category, error) {
	q := `UPDATE categories set name=$1 WHERE id=$2 AND user_id=$3`

	_, err := us.db.Exec(q, cat.Name, cat.ID, cat.UserID)
	if err != nil {
		return nil, err
	}

	return cat, nil
}
func (us *categoryStore) Get(cat *model.Category) (*model.Category, error) {
	q := `SELECT name FROM categories WHERE id=$1 AND user_id=$2`
	err := us.db.QueryRow(q, cat.ID, cat.UserID).Scan(&cat.Name)
	if err != nil {
		return nil, err
	}

	return cat, nil
}
func (us *categoryStore) Delete(cat *model.Category) error {
	q := `DELETE FROM categories WHERE id=$1 AND user_id=$2`

	_, err := us.db.Exec(q, cat.ID, cat.UserID)

	return err
}
