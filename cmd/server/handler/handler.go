package handler

import (
	"net/http"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/labstack/echo/v4"
)

func ErrHandler(err error, ctx echo.Context) {
	var er *errors.Error

	switch e := err.(type) {
	case errors.Error:
		er = &e
	case *errors.Error:
		er = e
	}

	if er == nil {
		ctx.JSON(http.StatusInternalServerError, nil)

		return
	}

	r := Response{
		Errors: er,
	}

	ctx.JSON(er.StatusCode, r)
}

func Handler(h func(echo.Context) (any, error)) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		data, err := h(ctx)
		if err != nil {
			return err
		}

		r := Response{
			Data: data,
		}

		return ctx.JSON(getCode(ctx.Request().Method), r)
	}
}

func getCode(method string) int {
	code := map[string]int{
		http.MethodGet:    http.StatusOK,
		http.MethodPost:   http.StatusCreated,
		http.MethodDelete: http.StatusNoContent,
		http.MethodPut:    http.StatusOK,
	}

	return code[method]
}
