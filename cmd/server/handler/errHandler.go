package handler

import (
	"net/http"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/labstack/echo/v4"
)

func ErrHandler(err error, c echo.Context) {
	var er *errors.Error

	switch e := err.(type) {
	case errors.Error:
		er = &e
	case *errors.Error:
		er = e
	}

	if er == nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// TODO: Complete the implementation
}
