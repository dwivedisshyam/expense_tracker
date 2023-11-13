package handler

import (
	"net/http"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/utils"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userSvc service.User
}

func NewUser(s service.User) userHandler {
	return userHandler{userSvc: s}
}

func (us *userHandler) Create(ctx echo.Context) error {
	var u model.User

	if err := ctx.Bind(&u); err != nil {
		return err
	}

	user, err := us.userSvc.Create(&u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (us *userHandler) Get(ctx echo.Context) error {
	id, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err
	}

	user, err := us.userSvc.Get(&model.User{ID: int64(id)})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (us *userHandler) Update(ctx echo.Context) error {
	var u model.User

	id, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err
	}

	if err := ctx.Bind(&u); err != nil {
		return err
	}

	u.ID = int64(id)

	user, err := us.userSvc.Update(&u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (us *userHandler) Delete(ctx echo.Context) error {
	id, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err
	}

	err = us.userSvc.Delete(&model.User{ID: id})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

func (us *userHandler) Login(ctx echo.Context) error {
	var u model.User

	if err := ctx.Bind(&u); err != nil {
		return err

	}

	token, err := us.userSvc.Login(&u)
	if err != nil {
		return err

	}

	return ctx.JSON(http.StatusOK, token)
}
