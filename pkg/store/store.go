package store

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/store/mongo"
	"github.com/dwivedisshyam/expense_tracker/pkg/store/sqlite"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/migration"
)

type User interface {
	Create(ctx *gofr.Context, user *model.User) (*model.User, error)
	Update(ctx *gofr.Context, user *model.User) error
	Get(ctx *gofr.Context, f *model.UserFilter) (*model.User, error)
	Delete(ctx *gofr.Context, f *model.UserFilter) error
}

type Category interface {
	Index(ctx *gofr.Context, f *model.CategoryFilter) ([]model.Category, error)
	Create(ctx *gofr.Context, cat *model.Category) (*model.Category, error)
	Update(ctx *gofr.Context, cat *model.Category) error
	Get(ctx *gofr.Context, f *model.CategoryFilter) (*model.Category, error)
	Delete(ctx *gofr.Context, f *model.CategoryFilter) error
}

type Expense interface {
	Index(ctx *gofr.Context, f *model.ExpenseFilter) ([]model.Expense, error)
	Create(ctx *gofr.Context, exp *model.Expense) (*model.Expense, error)
	Update(ctx *gofr.Context, exp *model.Expense) error
	Get(ctx *gofr.Context, f *model.ExpenseFilter) (*model.Expense, error)
	Delete(ctx *gofr.Context, f *model.ExpenseFilter) error
}

type Income interface {
	Index(ctx *gofr.Context, f *model.IncomeFilter) ([]model.Income, error)
	Create(ctx *gofr.Context, inc *model.Income) (*model.Income, error)
	Update(ctx *gofr.Context, inc *model.Income) error
	Get(ctx *gofr.Context, f *model.IncomeFilter) (*model.Income, error)
	Delete(ctx *gofr.Context, f *model.Income) error
}

func NewUser(app *gofr.App) User {
	if app.Config.Get("DB_TYPE") == "mongo" {
		return mongo.NewUser(CalculateNewID)
	}

	return sqlite.NewUser(CalculateNewID)
}

func NewCategory(app *gofr.App) Category {
	if app.Config.Get("DB_TYPE") == "mongo" {
		return mongo.NewCategory(CalculateNewID)
	}

	return sqlite.NewCategory(CalculateNewID)
}

func NewExpense(app *gofr.App) Expense {
	if app.Config.Get("DB_TYPE") == "mongo" {
		return mongo.NewExpense(CalculateNewID)
	}

	return sqlite.NewExpense(CalculateNewID)
}

// func NewIncome(cfg config.Config) Income {
// 	if cfg.Get("DB_TYPE") == "mongo" {
// 		return mongo.NewIncome(CalculateNewID)
// 	}

// 	return sqlite.NewIncome(CalculateNewID)
// }

func CalculateNewID(now time.Time) string {
	hash := md5.Sum([]byte(now.Format(time.RFC3339Nano)))
	return hex.EncodeToString(hash[:])[:7]
}

func Migrations() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{
		20240726091211: runMigration(),
	}
}

func runMigration() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(`CREATE TABLE IF NOT EXISTS users (
				id varchar(20) PRIMARY KEY,
				first_name varchar(20),
				last_name varchar(20),
				email varchar(25) unique,
				password varchar(255)
				)`)
			if err != nil {
				return err
			}

			_, err = d.SQL.Exec(`CREATE TABLE IF NOT EXISTS categories (
				id varchar(20) PRIMARY KEY, 
				name varchar(20) unique, 
				user_id varchar(20) REFERENCES users(id) ON DELETE CASCADE DEFAULT NULL
				)`)
			if err != nil {
				return err
			}

			_, err = d.SQL.Exec(`CREATE TABLE IF NOT EXISTS expenses (
				id varchar(20) PRIMARY KEY,
				user_id varchar(20) REFERENCES users(id) ON DELETE CASCADE, 
				category_id varchar(20) REFERENCES categories(id) ON DELETE CASCADE NOT NULL, 
				title varchar(20), 
				amount real, 
				due_date varchar(10),
				is_paid boolean
				)`)
			if err != nil {
				return err
			}

			_, err = d.SQL.Exec(`CREATE TABLE IF NOT EXISTS incomes (
				id varchar(20) PRIMARY KEY,
				user_id varchar(20) REFERENCES users(id) ON DELETE CASCADE, 
				title varchar(20), 
				amount real, 
				date varchar(10))`)
			if err != nil {
				return err
			}

			return nil
		},
	}

}
