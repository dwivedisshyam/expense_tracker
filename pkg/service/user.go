package service

import (
	"log"
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
	user.Password = ""
	return user, err
}
func (us *userSvc) Delete(user *model.User) error {
	return us.store.Delete(user)
}

func (us *userSvc) Login(user *model.User) (string, error) {
	var err error
	user, err = us.store.Get(&model.UserFilter{Email: user.Email})

	key := []byte(os.Getenv("JWT_KEY"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":     user.ID,
			"f_name": user.FName,
			"l_name": user.LName,
			"email":  user.Email,
		})

	s, err := t.SignedString(key)
	if err != nil {
		log.Println(err)
	}
	return s, err
}
