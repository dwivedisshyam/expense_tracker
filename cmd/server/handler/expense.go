package handler

import (
	"net/http"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/utils"
	"github.com/labstack/echo/v4"
)

type expHandler struct {
	expSvc service.Expense
}

func NewExpense(s service.Expense) expHandler {
	return expHandler{expSvc: s}
}

func (us *expHandler) Index(ctx echo.Context) error {
	var c model.ExpFilter

	userid, err := utils.ToInt64(ctx.Param("use_id"))
	if err != nil {
		return err
	}

	c.StartDate = ctx.Param("start_date")
	c.EndDate = ctx.Param("end_date")

	c.UserID = userid

	users, err := us.expSvc.Index(&c)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, users)
}

func (us *expHandler) Create(ctx echo.Context) error {
	var c model.Expense

	if err := ctx.Bind(&c); err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("use_id"))
	if err != nil {
		return err
	}

	c.UserID = userid

	user, err := us.expSvc.Create(&c)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (us *expHandler) Get(ctx echo.Context) error {
	id, err := utils.ToInt64(ctx.Param("id"))
	if err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("use_id"))
	if err != nil {
		return err
	}

	user, err := us.expSvc.Get(&model.Expense{ID: int64(id), UserID: int64(userid)})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (us *expHandler) Update(ctx echo.Context) error {
	var c model.Expense

	id, err := utils.ToInt64(ctx.Param("id"))
	if err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("use_id"))
	if err != nil {
		return err
	}

	if err := ctx.Bind(&c); err != nil {
		return err
	}

	c.ID = id
	c.UserID = userid

	user, err := us.expSvc.Update(&c)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (us *expHandler) Delete(ctx echo.Context) error {
	id, err := utils.ToInt64(ctx.Param("id"))
	if err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("use_id"))
	if err != nil {
		return err
	}

	err = us.expSvc.Delete(&model.Expense{ID: int64(id), UserID: int64(userid)})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
