package handler

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"gofr.dev/pkg/gofr"
)

type catHandler struct {
	catSvc service.Category
}

func NewCategory(s service.Category) catHandler {
	return catHandler{catSvc: s}
}

func (us *catHandler) Index(ctx *gofr.Context) (any, error) {
	var f model.CategoryFilter

	userid := ctx.PathParam("user_id")

	f.UserID = userid

	cats, err := us.catSvc.Index(ctx, &f)
	if err != nil {
		return nil, err
	}

	return cats, nil
}

func (us *catHandler) Create(ctx *gofr.Context) (any, error) {
	var c model.Category

	if err := ctx.Bind(&c); err != nil {
		return nil, err
	}

	userid := ctx.PathParam("user_id")

	c.UserID = userid

	newCat, err := us.catSvc.Create(ctx, &c)
	if err != nil {
		return nil, err
	}

	return newCat, nil
}

func (us *catHandler) Get(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")
	userid := ctx.PathParam("user_id")

	cat, err := us.catSvc.Get(ctx, &model.CategoryFilter{ID: id, UserID: userid})
	if err != nil {
		return nil, err
	}

	return cat, nil
}

func (us *catHandler) Update(ctx *gofr.Context) (any, error) {
	var c model.Category

	id := ctx.PathParam("id")
	userid := ctx.PathParam("user_id")
	if err := ctx.Bind(&c); err != nil {
		return nil, err
	}

	c.ID = id
	c.UserID = userid

	err := us.catSvc.Update(ctx, &c)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *catHandler) Delete(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")

	userid := ctx.PathParam("user_id")

	err := us.catSvc.Delete(ctx, &model.CategoryFilter{
		ID:     id,
		UserID: userid,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
