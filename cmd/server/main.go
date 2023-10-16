package main

import (
	"log"
	"net/http"

	"github.com/dwivedisshyam/expense_tracker/cmd/server/handler"
	"github.com/dwivedisshyam/expense_tracker/db"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	db := db.New()

	userStore := store.NewUser(db)
	expStore := store.NewExpense(db)
	catStore := store.NewCategory(db)
	incomeStore := store.NewIncome(db)

	userSvc := service.NewUser(userStore)
	catSvc := service.NewCategory(catStore)
	expSvc := service.NewExpense(expStore)
	incomeSvc := service.NewIncome(incomeStore)

	userH := handler.NewUser(userSvc)
	catH := handler.NewCategory(catSvc)
	expH := handler.NewExpense(expSvc)
	incH := handler.NewIncome(incomeSvc)

	r.HandleFunc("/login", userH.Login).Methods(http.MethodGet)

	r.HandleFunc("/users", userH.Create).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", userH.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", userH.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", userH.Delete).Methods(http.MethodDelete)

	r.HandleFunc("/users/{user_id}/categories", catH.Create).Methods(http.MethodPost)
	r.HandleFunc("/users/{user_id}/categories", catH.Index).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}/categories/{id}", catH.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}/categories/{id}", catH.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{user_id}/categories/{id}", catH.Delete).Methods(http.MethodDelete)

	r.HandleFunc("/users/{user_id}/expenses", expH.Create).Methods(http.MethodPost)
	r.HandleFunc("/users/{user_id}/expenses", expH.Index).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}/expenses/{id}", expH.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}/expenses/{id}", expH.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{user_id}/expenses/{id}", expH.Delete).Methods(http.MethodDelete)

	r.HandleFunc("/users/{user_id}/incomes", incH.Create).Methods(http.MethodPost)
	r.HandleFunc("/users/{user_id}/incomes/{id}", incH.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}/incomes/{id}", incH.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{user_id}/incomes/{id}", incH.Delete).Methods(http.MethodDelete)

	//fs := http.FileServer(http.Dir("./web/assets/"))
	//r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	//
	//r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	t, _ := template.ParseFiles("./web/pages-login.html")
	//	t.Execute(w, nil)
	//})
	//
	//r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
	//	t, _ := template.ParseFiles("./web/pages-register.html")
	//	t.Execute(w, nil)
	//})

	log.Println("Starting HTTP server @ :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
