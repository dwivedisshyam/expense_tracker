package main

import (
	"github.com/dwivedisshyam/expense_tracker/cmd/server/handler"
	"github.com/dwivedisshyam/expense_tracker/db"
	"github.com/dwivedisshyam/expense_tracker/pkg"
	"github.com/dwivedisshyam/expense_tracker/pkg/middleware"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
)

func main() {
	app := pkg.NewApp()

	app.Use(middleware.Auth(app))

	db := db.New()

	userStore := store.NewUser(db)
	expStore := store.NewExpense(db)
	catStore := store.NewCategory(db)
	incomeStore := store.NewIncome(db)

	userSvc := service.NewUser(app, userStore)
	catSvc := service.NewCategory(catStore)
	expSvc := service.NewExpense(expStore)
	incomeSvc := service.NewIncome(incomeStore)

	userH := handler.NewUser(userSvc)
	catH := handler.NewCategory(catSvc)
	expH := handler.NewExpense(expSvc)
	incH := handler.NewIncome(incomeSvc)

	app.POST("/login", userH.Login)

	app.POST("/users", userH.Create)
	app.GET("/users/{user_id}", userH.Get)
	app.PUT("/users/{user_id}", userH.Update)
	app.DELETE("/users/{user_id}", userH.Delete)

	app.POST("/users/{user_id}/categories", catH.Create)
	app.GET("/users/{user_id}/categories", catH.Index)
	app.GET("/users/{user_id}/categories/{id}", catH.Get)
	app.PUT("/users/{user_id}/categories/{id}", catH.Update)
	app.DELETE("/users/{user_id}/categories/{id}", catH.Delete)

	app.POST("/users/{user_id}/expenses", expH.Create)
	app.GET("/users/{user_id}/expenses", expH.Index)
	app.GET("/users/{user_id}/expenses/{id}", expH.Get)
	app.PUT("/users/{user_id}/expenses/{id}", expH.Update)
	app.DELETE("/users/{user_id}/expenses/{id}", expH.Delete)

	app.POST("/users/{user_id}/incomes", incH.Create)
	app.GET("/users/{user_id}/incomes/{id}", incH.Get)
	app.PUT("/users/{user_id}/incomes/{id}", incH.Update)
	app.DELETE("/users/{user_id}/incomes/{id}", incH.Delete)

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

	app.Run()
}
