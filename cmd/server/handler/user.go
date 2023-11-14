package handler

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/utils"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userSvc service.User
}

func NewUser(s service.User) userHandler {
	return userHandler{userSvc: s}
}

func (us *userHandler) Create(ctx echo.Context) (any, error) {
	var u model.User

	if err := ctx.Bind(&u); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	user, err := us.userSvc.Create(&u)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userHandler) Get(ctx echo.Context) (any, error) {
	id, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	user, err := us.userSvc.Get(&model.User{ID: id})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userHandler) Update(ctx echo.Context) (any, error) {
	var u model.User

	id, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	if err := ctx.Bind(&u); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	u.ID = id

	user, err := us.userSvc.Update(&u)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userHandler) Delete(ctx echo.Context) (any, error) {
	id, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	err = us.userSvc.Delete(&model.User{ID: id})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *userHandler) Login(ctx echo.Context) (any, error) {
	var u model.User

	if err := ctx.Bind(&u); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	token, err := us.userSvc.Login(&u)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token": token,
	}, nil
}
