package handler

import (
	"net/http"
	"strconv"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/gorilla/mux"
)

type incHandler struct {
	incSvc service.Income
}

func NewIncome(s service.Income) incHandler {
	return incHandler{incSvc: s}
}

func (us *incHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c model.Income

	resp := Responder{w}

	if err := Bind(r, &c); err != nil {
		resp.Respond(nil, err)
		return
	}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	c.UserID = int64(userid)

	user, err := us.incSvc.Create(&c)
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *incHandler) Get(w http.ResponseWriter, r *http.Request) {
	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	user, err := us.incSvc.Get(&model.Income{ID: int64(id), UserID: int64(userid)})
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *incHandler) Update(w http.ResponseWriter, r *http.Request) {
	var c model.Income

	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}
	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	if err := Bind(r, &c); err != nil {
		resp.Respond(nil, err)
		return
	}

	c.ID = int64(id)
	c.UserID = int64(userid)

	user, err := us.incSvc.Update(&c)
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *incHandler) Delete(w http.ResponseWriter, r *http.Request) {
	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	err = us.incSvc.Delete(&model.Income{ID: int64(id), UserID: int64(userid)})
	if err != nil {
		resp.Respond(nil, err)
		return
	}
}
