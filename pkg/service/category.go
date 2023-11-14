package service

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
)

type categorySvc struct {
	store store.Category
}

func NewCategory(s store.Category) Category {
	return &categorySvc{store: s}
}

func (us *categorySvc) Index(f *model.CatFilter) ([]model.Category, error) {
	return us.store.Index(f)
}

func (us *categorySvc) Create(cat *model.Category) (*model.Category, error) {
	return us.store.Create(cat)
}

func (us *categorySvc) Update(cat *model.Category) (*model.Category, error) {
	return us.store.Update(cat)
}

func (us *categorySvc) Get(cat *model.Category) (*model.Category, error) {
	return us.store.Get(cat)
}

func (us *categorySvc) Delete(cat *model.Category) error {
	return us.store.Delete(cat)
}
