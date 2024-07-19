package service

import (
	"crypto/sha512"
	"fmt"
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"gofr.dev/pkg/gofr"
)

const (
	day = time.Hour * 24
)

type userSvc struct {
	store  store.User
	jwtKey []byte
}

func NewUser(jwtKey string, s store.User) User {
	return &userSvc{
		jwtKey: []byte(jwtKey),
		store:  s,
	}
}

func hashPassowrd(password string) string {
	hash := sha512.Sum512([]byte(password))

	return fmt.Sprintf("%x", hash)
}

func (us *userSvc) Create(ctx *gofr.Context, user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	user.Password = hashPassowrd(user.Password)

	err := us.store.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (us *userSvc) Update(ctx *gofr.Context, user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	user.Password = hashPassowrd(user.Password)

	return us.store.Update(ctx, user)
}

func (us *userSvc) Get(ctx *gofr.Context, filter *model.UserFilter) (*model.User, error) {
	user, err := us.store.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (us *userSvc) Delete(ctx *gofr.Context, filter *model.UserFilter) error {
	return us.store.Delete(ctx, filter)
}

func (us *userSvc) Login(ctx *gofr.Context, user *model.User) (string, error) {
	u, err := us.store.Get(ctx, &model.UserFilter{Email: user.Email})
	if err != nil {
		ctx.Logger.Error(err)
		return "", errors.Unauthenticated("invalid credentials")
	}

	if !matchPassword(user.Password, u.Password) {
		return "", errors.Unauthenticated("invalid credentials")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		model.Claims{
			ID:    u.ID,
			Email: u.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(day)),
			},
		})

	s, err := t.SignedString(us.jwtKey)
	if err != nil {
		return "", err
	}

	return s, err
}

func matchPassword(paswd, hash string) bool {
	hpswd := hashPassowrd(paswd)

	return hpswd == hash
}
