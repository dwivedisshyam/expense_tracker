package service

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
	"gofr.dev/pkg/gofr"
)

type categorySvc struct {
	store store.Category
}

func NewCategory(store store.Category) Category {
	return &categorySvc{store: store}
}

func (us *categorySvc) Index(ctx *gofr.Context, f *model.CategoryFilter) ([]model.Category, error) {
	return us.store.Index(ctx, f)
}

func (us *categorySvc) Create(ctx *gofr.Context, cat *model.Category) (*model.Category, error) {
	return us.store.Create(ctx, cat)
}

func (us *categorySvc) Update(ctx *gofr.Context, cat *model.Category) error {
	return us.store.Update(ctx, cat)
}

func (us *categorySvc) Get(ctx *gofr.Context, filter *model.CategoryFilter) (*model.Category, error) {
	return us.store.Get(ctx, filter)
}

func (us *categorySvc) Delete(ctx *gofr.Context, cat *model.CategoryFilter) error {
	return us.store.Delete(ctx, cat)
}
