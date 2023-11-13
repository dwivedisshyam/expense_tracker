package handler

import (
	"net/http"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/labstack/echo/v4"
)

func Respond(ctx echo.Context, data any, err error) {
	status := http.StatusOK

	r := &Response{
		Data:   data,
		Errors: err,
	}

	if err != nil {
		status = http.StatusInternalServerError

		if er, ok := err.(*errors.Error); ok {
			status = er.StatusCode
		}
	}

	ctx.JSON(status, r)
}

type Response struct {
	Data   any   `json:"data,omitempty"`
	Errors error `json:"errors,omitempty"`
}
