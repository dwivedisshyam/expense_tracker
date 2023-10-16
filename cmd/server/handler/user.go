package handler

import (
	"net/http"
	"strconv"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/gorilla/mux"
)

type userHandler struct {
	userSvc service.User
}

func NewUser(s service.User) userHandler {
	return userHandler{userSvc: s}
}

func (us *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u model.User

	resp := Responder{w}

	if err := Bind(r, &u); err != nil {
		resp.Respond(nil, err)
		return
	}

	user, err := us.userSvc.Create(&u)
	if err != nil {
		resp.Respond(nil, err)
		return
	}
	resp.Respond(user, nil)
}

func (us *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	user, err := us.userSvc.Get(&model.User{ID: int64(id)})
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	var u model.User

	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	if err := Bind(r, &u); err != nil {
		resp.Respond(nil, err)
		return
	}

	u.ID = int64(id)

	user, err := us.userSvc.Update(&u)
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	err = us.userSvc.Delete(&model.User{ID: int64(id)})
	if err != nil {
		resp.Respond(nil, err)
		return
	}
}

func (us *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var u model.User

	resp := Responder{w}

	if err := Bind(r, &u); err != nil {
		resp.Respond(nil, err)
		return
	}

	token, err := us.userSvc.Login(&u)
	if err != nil {
		resp.Respond(nil, err)
		return
	}
	resp.Respond(token, nil)
}
