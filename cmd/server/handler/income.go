package handler

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"gofr.dev/pkg/gofr"
)

type incHandler struct {
	incSvc service.Income
}

func NewIncome(s service.Income) incHandler {
	return incHandler{incSvc: s}
}

func (us *incHandler) Create(ctx *gofr.Context) (any, error) {
	var c model.Income

	if err := ctx.Bind(&c); err != nil {
		return nil, err
	}

	c.UserID = ctx.Param("user_id")

	err := us.incSvc.Create(&c)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *incHandler) Get(ctx *gofr.Context) (any, error) {
	id := ctx.Param("id")
	userid := ctx.Param("user_id")

	income, err := us.incSvc.Get(&model.Income{ID: id, UserID: userid})
	if err != nil {
		return nil, err
	}

	return income, nil
}

func (us *incHandler) Update(ctx *gofr.Context) (any, error) {
	var c model.Income

	c.ID = ctx.Param("id")
	c.UserID = ctx.Param("user_id")

	if err := ctx.Bind(&c); err != nil {
		return nil, err
	}

	err := us.incSvc.Update(&c)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *incHandler) Delete(ctx *gofr.Context) (any, error) {
	id := ctx.Param("id")
	userid := ctx.Param("user_id")

	err := us.incSvc.Delete(&model.Income{ID: id, UserID: userid})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
