package handler

import (
	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type userHandler struct {
	userSvc service.User
}

func NewUser(s service.User) userHandler {
	return userHandler{userSvc: s}
}

func (us *userHandler) Create(ctx *gofr.Context) (any, error) {
	var u model.User

	if err := ctx.Bind(&u); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	err := us.userSvc.Create(ctx, &u)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *userHandler) Get(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("user_id")

	user, err := us.userSvc.Get(ctx, &model.UserFilter{ID: id})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userHandler) Update(ctx *gofr.Context) (any, error) {
	var u model.User

	id := ctx.PathParam("user_id")

	if err := ctx.Bind(&u); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	u.ID = id

	err := us.userSvc.Update(ctx, &u)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *userHandler) Delete(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("user_id")
	err := us.userSvc.Delete(ctx, &model.UserFilter{ID: id})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (us *userHandler) Login(ctx *gofr.Context) (any, error) {
	var u model.User

	if err := ctx.Bind(&u); err != nil {
		return nil, errors.BadRequest(err.Error())
	}

	token, err := us.userSvc.Login(ctx, &u)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"token": token,
	}, nil
}
