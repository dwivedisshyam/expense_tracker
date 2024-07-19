package handler

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"gofr.dev/pkg/gofr"
)

type expHandler struct {
	expSvc service.Expense
}

func NewExpense(s service.Expense) expHandler {
	return expHandler{expSvc: s}
}

func (us *expHandler) Index(ctx *gofr.Context) (any, error) {
	var c model.ExpenseFilter

	userID := ctx.PathParam("user_id")

	c.StartDate = ctx.Param("start_date")
	c.EndDate = ctx.Param("end_date")

	c.UserID = userID

	exps, err := us.expSvc.Index(ctx, &c)
	if err != nil {
		return nil, err
	}

	return exps, nil
}

func (us *expHandler) Create(ctx *gofr.Context) (any, error) {
	var exp model.Expense

	if err := ctx.Bind(&exp); err != nil {
		return nil, err
	}

	userID := ctx.PathParam("user_id")

	exp.UserID = userID

	newExp, err := us.expSvc.Create(ctx, &exp)
	if err != nil {
		return nil, err
	}

	return newExp, nil
}

func (us *expHandler) Get(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")
	userid := ctx.PathParam("user_id")

	exp, err := us.expSvc.Get(ctx, &model.ExpenseFilter{ID: id, UserID: userid})
	if err != nil {
		return nil, err
	}

	return exp, nil
}

func (us *expHandler) Update(ctx *gofr.Context) (any, error) {
	var exp model.Expense

	id := ctx.PathParam("id")
	userid := ctx.PathParam("user_id")

	if err := ctx.Bind(&exp); err != nil {
		return nil, err
	}

	exp.ID = id
	exp.UserID = userid

	err := us.expSvc.Update(ctx, &exp)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *expHandler) Delete(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")
	userid := ctx.PathParam("user_id")

	err := us.expSvc.Delete(ctx, &model.ExpenseFilter{ID: id, UserID: userid})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
