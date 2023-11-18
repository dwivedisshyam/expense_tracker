package main

import (
	"github.com/dwivedisshyam/expense_tracker/cmd/server/handler"
	"github.com/dwivedisshyam/expense_tracker/db"
	"github.com/dwivedisshyam/expense_tracker/pkg/middleware"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// nolint: funlen
func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.HTTPErrorHandler = handler.ErrHandler

	e.Use(middleware.Auth)

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

	e.POST("/login", handler.Handler(userH.Login))

	e.POST("/users", handler.Handler(userH.Create))
	e.GET("/users/:user_id", handler.Handler(userH.Get))
	e.PUT("/users/:user_id", handler.Handler(userH.Update))
	e.DELETE("/users/:user_id", handler.Handler(userH.Delete))

	e.POST("/users/:user_id/categories", catH.Create)
	e.GET("/users/:user_id/categories", catH.Index)
	e.GET("/users/:user_id/categories/:id", catH.Get)
	e.PUT("/users/:user_id/categories/:id", catH.Update)
	e.DELETE("/users/:user_id/categories/:id", catH.Delete)

	e.POST("/users/:user_id/expenses", expH.Create)
	e.GET("/users/:user_id/expenses", expH.Index)
	e.GET("/users/:user_id/expenses/:id", expH.Get)
	e.PUT("/users/:user_id/expenses/:id", expH.Update)
	e.DELETE("/users/:user_id/expenses/:id", expH.Delete)

	e.POST("/users/:user_id/incomes", incH.Create)
	e.GET("/users/:user_id/incomes/:id", incH.Get)
	e.PUT("/users/:user_id/incomes/:id", incH.Update)
	e.DELETE("/users/:user_id/incomes/:id", incH.Delete)

	// fs := http.FileServer(http.Dir("./web/assets/"))
	// r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	// r.HandleFunc("/", func(ctx echo.Context) {
	// 	t, _ := template.ParseFiles("./web/pages-login.html")
	// 	t.Execute(w, nil)
	// })

	// r.HandleFunc("/register", func(ctx echo.Context) {
	// 	t, _ := template.ParseFiles("./web/pages-register.html")
	// 	t.Execute(w, nil)
	// })

	e.Logger.Fatal(e.Start(":8000"))
}
