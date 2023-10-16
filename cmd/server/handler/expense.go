package handler

import (
	"net/http"
	"strconv"

	"github.com/dwivedisshyam/expense_tracker/pkg/model"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/gorilla/mux"
)

type expHandler struct {
	expSvc service.Expense
}

func NewExpense(s service.Expense) expHandler {
	return expHandler{expSvc: s}
}

func (us *expHandler) Index(w http.ResponseWriter, r *http.Request) {
	var c model.ExpFilter

	resp := Responder{w}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	c.StartDate = r.URL.Query().Get("start_date")
	c.EndDate = r.URL.Query().Get("end_date")

	c.UserID = int64(userid)

	user, err := us.expSvc.Index(&c)
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *expHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c model.Expense

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

	user, err := us.expSvc.Create(&c)
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *expHandler) Get(w http.ResponseWriter, r *http.Request) {
	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	user, err := us.expSvc.Get(&model.Expense{ID: int64(id), UserID: int64(userid)})
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *expHandler) Update(w http.ResponseWriter, r *http.Request) {
	var c model.Expense

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

	user, err := us.expSvc.Update(&c)
	if err != nil {
		resp.Respond(nil, err)
		return
	}

	resp.Respond(user, nil)
}

func (us *expHandler) Delete(w http.ResponseWriter, r *http.Request) {
	resp := Responder{w}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	userid, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		resp.Respond(nil, err)
	}

	err = us.expSvc.Delete(&model.Expense{ID: int64(id), UserID: int64(userid)})
	if err != nil {
		resp.Respond(nil, err)
		return
	}
}
