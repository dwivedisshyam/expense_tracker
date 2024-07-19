package service

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"gofr.dev/pkg/gofr"
)

type User interface {
	Login(ctx *gofr.Context, user *model.User) (string, error)

	Create(ctx *gofr.Context, user *model.User) (*model.User, error)
	Update(ctx *gofr.Context, user *model.User) error
	Get(ctx *gofr.Context, user *model.UserFilter) (*model.User, error)
	Delete(ctx *gofr.Context, user *model.UserFilter) error
}

type Category interface {
	Index(ctx *gofr.Context, f *model.CategoryFilter) ([]model.Category, error)
	Create(ctx *gofr.Context, user *model.Category) (*model.Category, error)
	Update(ctx *gofr.Context, user *model.Category) error
	Get(ctx *gofr.Context, user *model.CategoryFilter) (*model.Category, error)
	Delete(ctx *gofr.Context, user *model.CategoryFilter) error
}

type Expense interface {
	Index(ctx *gofr.Context, f *model.ExpenseFilter) ([]model.Expense, error)
	Create(ctx *gofr.Context, exp *model.Expense) (*model.Expense, error)
	Update(ctx *gofr.Context, exp *model.Expense) error
	Get(ctx *gofr.Context, filter *model.ExpenseFilter) (*model.Expense, error)
	Delete(ctx *gofr.Context, filter *model.ExpenseFilter) error
}

type Income interface {
	Create(user *model.Income) error
	Update(user *model.Income) error
	Get(user *model.Income) (*model.Income, error)
	Delete(user *model.Income) error
}
