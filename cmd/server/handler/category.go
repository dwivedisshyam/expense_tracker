package handler

import (
	"net/http"
	"strconv"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/utils"
	"github.com/gorilla/mux"
)

type catHandler struct {
	catSvc service.Category
}

func NewCategory(s service.Category) catHandler {
	return catHandler{catSvc: s}
}

func (us *catHandler) Index(w http.ResponseWriter, r *http.Request) {
	var f model.CatFilter

	resp := Responder{w}
	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	f.UserID = int64(userid)
	cats, err := us.catSvc.Index(&f)
	if err != nil {
		resp.Respond(nil, err)
	}

	resp.Respond(cats, nil)

}

func (us *catHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c model.Category

	resp := Responder{w}

	if err := utils.Bind(r, &c); err != nil {
		resp.Respond(nil, err)
		return
	}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	c.UserID = int64(userid)

	cat, err := us.catSvc.Create(&c)
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(cat, nil)
}

func (us *catHandler) Get(w http.ResponseWriter, r *http.Request) {
	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	user, err := us.catSvc.Get(&model.Category{ID: int64(id), UserID: int64(userid)})
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *catHandler) Update(w http.ResponseWriter, r *http.Request) {
	var c model.Category

	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	if err := utils.Bind(r, &c); err != nil {
		resp.Respond(nil, err)
		return
	}

	c.ID = int64(id)
	c.UserID = int64(userid)

	user, err := us.catSvc.Update(&c)
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *catHandler) Delete(w http.ResponseWriter, r *http.Request) {
	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	err = us.catSvc.Delete(&model.Category{
		ID:     int64(id),
		UserID: int64(userid),
	})
	if err != nil {
		resp.Respond(nil, err)
		return
	}
}
