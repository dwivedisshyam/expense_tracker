package handler

import (
	"net/http"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/utils"
	"github.com/labstack/echo/v4"
)

type catHandler struct {
	catSvc service.Category
}

func NewCategory(s service.Category) catHandler {
	return catHandler{catSvc: s}
}

func (us *catHandler) Index(ctx echo.Context) error {
	var f model.CatFilter

	userid, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err

	}

	f.UserID = userid
	cats, err := us.catSvc.Index(&f)
	if err != nil {
		return err

	}

	return ctx.JSON(http.StatusOK, cats)
}

func (us *catHandler) Create(ctx echo.Context) error {
	var c model.Category

	if err := ctx.Bind(&c); err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err
	}

	c.UserID = int64(userid)

	cat, err := us.catSvc.Create(&c)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, cat)
}

func (us *catHandler) Get(ctx echo.Context) error {
	id, err := utils.ToInt64(ctx.Param("id"))
	if err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err
	}

	user, err := us.catSvc.Get(&model.Category{ID: int64(id), UserID: int64(userid)})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (us *catHandler) Update(ctx echo.Context) error {
	var c model.Category

	id, err := utils.ToInt64(ctx.Param("id"))
	if err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err
	}

	if err := ctx.Bind(&c); err != nil {
		return err
	}

	c.ID = id
	c.UserID = userid

	user, err := us.catSvc.Update(&c)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

func (us *catHandler) Delete(ctx echo.Context) error {
	id, err := utils.ToInt64(ctx.Param("id"))
	if err != nil {
		return err
	}

	userid, err := utils.ToInt64(ctx.Param("user_id"))
	if err != nil {
		return err
	}

	err = us.catSvc.Delete(&model.Category{
		ID:     int64(id),
		UserID: int64(userid),
	})

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
