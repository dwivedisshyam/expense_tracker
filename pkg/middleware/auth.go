package middleware

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func exemptPath(path string) bool {
	paths := map[string]bool{
		"POST /login": true,
		"POST /users": true,
	}
	_, ok := paths[path]

	return ok
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if exemptPath(ctx.Request().Method + " " + ctx.Path()) {
			return next(ctx)
		}

		authToken := ctx.Request().Header.Get("Authorization")
		if authToken == "" {
			return errors.Unauthenticated("invalid auth token")
		}

		authToken = strings.TrimPrefix(authToken, "Bearer ")

		key := []byte(os.Getenv("JWT_KEY"))
		if len(key) == 0 {
			ctx.Logger().Error("JWT_KEY is missing")

			return errors.Unauthenticated("invalid auth token")
		}

		claims := &model.Claims{}

		t, err := jwt.ParseWithClaims(authToken, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			ctx.Logger().Error(err)

			return errors.Unauthenticated("invalid auth token")
		}

		if !t.Valid {
			log.Println("invalid token")
			return errors.Unauthenticated("invalid auth token")
		}

		id := strconv.Itoa(int(claims.ID))

		if id != ctx.Param("user_id") {
			return errors.Unauthorized("un-authorized request")
		}

		return next(ctx)
	}
}
