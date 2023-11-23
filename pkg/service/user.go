package service

import (
	"crypto/sha512"
	"fmt"
	"os"
	"time"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
)

const (
	day = time.Hour * 24
)

type userSvc struct {
	store store.User
}

func NewUser(s store.User) User {
	return &userSvc{store: s}
}

func hashPassowrd(password string) string {
	hash := sha512.Sum512([]byte(password))

	return fmt.Sprintf("%x", hash)
}

func (us *userSvc) Create(user *model.User) (*model.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Password = hashPassowrd(user.Password)

	user, err := us.store.Create(user)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (us *userSvc) Update(user *model.User) (*model.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return us.store.Update(user)
}

func (us *userSvc) Get(user *model.User) (*model.User, error) {
	user, err := us.store.Get(&model.UserFilter{ID: user.ID})
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (us *userSvc) Delete(user *model.User) error {
	return us.store.Delete(user)
}

func (us *userSvc) Login(user *model.User) (string, error) {
	u, err := us.store.Get(&model.UserFilter{Email: user.Email})
	if err != nil {
		return "", err
	}

	err = validatePassword(user.Password, u.Password)
	if err != nil {
		return "", err
	}

	key := []byte(os.Getenv("JWT_KEY"))
	if len(key) == 0 {
		return "", errors.Unexpected("JWT_KEY missing")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		model.Claims{
			ID:    u.ID,
			Email: u.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(day)),
			},
		})

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, err
}

func validatePassword(paswd, hash string) error {
	hpswd := hashPassowrd(paswd)

	if hpswd != hash {
		return errors.Unauthenticated("invalid credentials")
	}

	return nil
}
