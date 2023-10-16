package service

import (
	"errors"
	"os"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
	"github.com/golang-jwt/jwt/v5"
)

type userSvc struct {
	store store.User
}

func NewUser(s store.User) User {
	return &userSvc{store: s}
}

func (us *userSvc) Create(user *model.User) (*model.User, error) {
	var err error
	user, err = us.store.Create(user)
	user.Password = ""
	return user, err
}
func (us *userSvc) Update(user *model.User) (*model.User, error) {
	return us.store.Update(user)
}
func (us *userSvc) Get(user *model.User) (*model.User, error) {
	var err error
	user, err = us.store.Get(&model.UserFilter{ID: user.ID})
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, err
}
func (us *userSvc) Delete(user *model.User) error {
	return us.store.Delete(user)
}

func (us *userSvc) Login(user *model.User) (string, error) {
	user, err := us.store.Get(&model.UserFilter{Email: user.Email})
	if err != nil {
		return "", err
	}

	key := []byte(os.Getenv("JWT_KEY"))
	if len(key) == 0 {
		return "", errors.New("JWT_KEY missing")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":     user.ID,
			"f_name": user.FName,
			"l_name": user.LName,
			"email":  user.Email,
		})

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, err
}
