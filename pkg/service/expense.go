package service

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
	"gofr.dev/pkg/gofr"
)

type expSvc struct {
	store store.Expense
}

func NewExpense(s store.Expense) Expense {
	return &expSvc{store: s}
}

func (us *expSvc) Index(ctx *gofr.Context, f *model.ExpenseFilter) ([]model.Expense, error) {
	return us.store.Index(ctx, f)
}

func (us *expSvc) Create(ctx *gofr.Context, exp *model.Expense) (*model.Expense, error) {
	if err := exp.Validate(); err != nil {
		return nil, err
	}

	return us.store.Create(ctx, exp)
}

func (us *expSvc) Update(ctx *gofr.Context, exp *model.Expense) error {
	if err := exp.Validate(); err != nil {
		return err
	}

	return us.store.Update(ctx, exp)
}

func (us *expSvc) Get(ctx *gofr.Context, filter *model.ExpenseFilter) (*model.Expense, error) {
	return us.store.Get(ctx, filter)
}

func (us *expSvc) Delete(ctx *gofr.Context, filter *model.ExpenseFilter) error {
	return us.store.Delete(ctx, filter)
}
