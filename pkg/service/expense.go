package service

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
)

type expSvc struct {
	store store.Expense
}

func NewExpense(s store.Expense) Expense {
	return &expSvc{store: s}
}

func (us *expSvc) Index(f *model.ExpFilter) ([]model.Expense, error) {
	return us.store.Index(f)
}

func (us *expSvc) Create(cat *model.Expense) (*model.Expense, error) {
	return us.store.Create(cat)
}

func (us *expSvc) Update(cat *model.Expense) (*model.Expense, error) {
	return us.store.Update(cat)
}

func (us *expSvc) Get(cat *model.Expense) (*model.Expense, error) {
	return us.store.Get(cat)
}

func (us *expSvc) Delete(cat *model.Expense) error {
	return us.store.Delete(cat)
}
