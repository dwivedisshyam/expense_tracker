package service

import "github.com/dwivedisshyam/expense_tracker/pkg/model"

type User interface {
	Login(user *model.User) (string, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Get(user *model.User) (*model.User, error)
	Delete(user *model.User) error
}

type Category interface {
	Index(f *model.CatFilter) ([]model.Category, error)
	Create(user *model.Category) (*model.Category, error)
	Update(user *model.Category) (*model.Category, error)
	Get(user *model.Category) (*model.Category, error)
	Delete(user *model.Category) error
}

type Expense interface {
	Index(f *model.ExpFilter) ([]model.Expense, error)
	Create(user *model.Expense) (*model.Expense, error)
	Update(user *model.Expense) (*model.Expense, error)
	Get(user *model.Expense) (*model.Expense, error)
	Delete(user *model.Expense) error
}

type Income interface {
	Create(user *model.Income) (*model.Income, error)
	Update(user *model.Income) (*model.Income, error)
	Get(user *model.Income) (*model.Income, error)
	Delete(user *model.Income) error
}
