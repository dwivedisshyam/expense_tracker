package store

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"gofr.dev/pkg/gofr"
)

const (
	CollectionUser     = "users"
	CollectionCategory = "categories"
	CollectionExpense  = "expenses"
)

type User interface {
	Create(ctx *gofr.Context, user *model.User) error
	Update(ctx *gofr.Context, user *model.User) error
	Get(ctx *gofr.Context, user *model.UserFilter) (*model.User, error)
	Delete(ctx *gofr.Context, user *model.UserFilter) error
}

type Category interface {
	Index(ctx *gofr.Context, f *model.CategoryFilter) ([]model.Category, error)
	Create(ctx *gofr.Context, user *model.Category) error
	Update(ctx *gofr.Context, user *model.Category) error
	Get(ctx *gofr.Context, user *model.CategoryFilter) (*model.Category, error)
	Delete(ctx *gofr.Context, user *model.Category) error
}

type Expense interface {
	Index(ctx *gofr.Context, f *model.ExpenseFilter) ([]model.Expense, error)
	Create(ctx *gofr.Context, exp *model.Expense) (*model.Expense, error)
	Update(ctx *gofr.Context, exp *model.Expense) error
	Get(ctx *gofr.Context, f *model.ExpenseFilter) (*model.Expense, error)
	Delete(ctx *gofr.Context, f *model.ExpenseFilter) error
}

type Income interface {
	Create(user *model.Income) error
	Update(user *model.Income) error
	Get(user *model.Income) (*model.Income, error)
	Delete(user *model.Income) error
}

func calculateNewID(now time.Time) string {
	hash := md5.Sum([]byte(now.Format(time.RFC3339Nano)))
	return hex.EncodeToString(hash[:])[:7]
}
