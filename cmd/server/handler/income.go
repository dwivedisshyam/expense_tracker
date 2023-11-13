package handler

import (
	"net/http"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/utils"
	"github.com/labstack/echo/v4"
)

type incHandler struct {
	incSvc service.Income
}

func NewIncome(s service.Income) incHandler {
	return incHandler{incSvc: s}
}

func (us *incHandler) Create(ctx echo.Context) error {
	var c model.Income

	if err := ctx.Bind(&c); err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err
	}

	c.UserID = int64(userid)

	user, err := us.incSvc.Create(&c)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (us *incHandler) Get(ctx echo.Context) error {
	id, err := utils.ToInt64(ctx.Param("id"))
	if err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err
	}

	user, err := us.incSvc.Get(&model.Income{ID: id, UserID: userid})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (us *incHandler) Update(ctx echo.Context) error {
	var c model.Income

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

	user, err := us.incSvc.Update(&c)
	if err != nil {
		return err

	}

	return ctx.JSON(http.StatusOK, user)
}

func (us *incHandler) Delete(ctx echo.Context) error {
	id, err := utils.ToInt64(ctx.Param("id"))
	if err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("use_id"))
	if err != nil {
		return err
	}

	err = us.incSvc.Delete(&model.Income{ID: int64(id), UserID: int64(userid)})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
