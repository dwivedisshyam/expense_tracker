package sqlite

import (
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"gofr.dev/pkg/gofr"
)

type storeIncome struct {
	idGen func(time.Time) string
}

func NewIncome(idGen func(time.Time) string) *storeIncome {
	return &storeIncome{idGen: idGen}
}

func (si storeIncome) Index(ctx *gofr.Context, f *model.IncomeFilter) ([]model.Income, error) {
	return nil, nil
}
func (si storeIncome) Create(ctx *gofr.Context, exp *model.Income) (*model.Income, error) {
	return nil, nil
}
func (si storeIncome) Update(ctx *gofr.Context, exp *model.Income) error {
	return nil
}
func (si storeIncome) Get(ctx *gofr.Context, f *model.IncomeFilter) (*model.Income, error) {
	return nil, nil
}
func (si storeIncome) Delete(ctx *gofr.Context, f *model.Income) error {
	return nil
}
