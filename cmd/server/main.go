package main

import (
	"strconv"

	"github.com/dwivedisshyam/expense_tracker/cmd/server/handler"
	"github.com/dwivedisshyam/expense_tracker/pkg/middleware"
	"github.com/dwivedisshyam/expense_tracker/pkg/service"
	"github.com/dwivedisshyam/expense_tracker/pkg/store"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource/mongo"
)

// nolint: funlen
func main() {
	app := gofr.New()

	if app.Config.Get("DB_TYPE") == "mongo" {
		dbPort, err := strconv.Atoi(app.Config.Get("DB_PORT"))
		if err != nil {
			app.Logger().Fatalf("missing env %q", "DB_PORT")
			return
		}
		mongoCfg := mongo.Config{
			Host:     app.Config.Get("DB_HOST"),
			User:     app.Config.Get("DB_USER"),
			Password: app.Config.Get("DB_PASSWORD"),
			Port:     dbPort,
			Database: app.Config.Get("DB_NAME"),
		}
		db := mongo.New(mongoCfg)
		app.AddMongo(db)
	} else {
		app.Migrate(store.Migrations())
	}

	jwtKey := app.Config.Get("JWT_KEY")
	if jwtKey == "" {
		app.Logger().Fatalf("missing env %q", "JWT_KEY")
		return
	}

	// Store
	userStore := store.NewUser(app)
	catStore := store.NewCategory(app)
	expStore := store.NewExpense(app)

	// Service
	userSvc := service.NewUser(jwtKey, userStore)
	catSvc := service.NewCategory(catStore)
	expSvc := service.NewExpense(expStore)

	// Handler
	userH := handler.NewUser(userSvc)
	catH := handler.NewCategory(catSvc)
	expH := handler.NewExpense(expSvc)

	app.POST("/login", userH.Login)

	app.POST("/users", userH.Create)
	app.GET("/users/{user_id}", userH.Get)
	app.PUT("/users/{user_id}", userH.Update)
	app.DELETE("/users/{user_id}", userH.Delete)

	app.GET("/users/{user_id}/categories", catH.Index)
	app.POST("/users/{user_id}/categories", catH.Create)
	app.GET("/users/{user_id}/categories/{id}", catH.Get)
	app.PUT("/users/{user_id}/categories/{id}", catH.Update)
	app.DELETE("/users/{user_id}/categories/{id}", catH.Delete)

	app.POST("/users/{user_id}/expenses", expH.Create)
	app.GET("/users/{user_id}/expenses", expH.Index)
	app.GET("/users/{user_id}/expenses/{id}", expH.Get)
	app.PUT("/users/{user_id}/expenses/{id}", expH.Update)
	app.DELETE("/users/{user_id}/expenses/{id}", expH.Delete)

	// middleware
	app.UseMiddleware(middleware.Authentication(jwtKey))

	app.Run()
}

// incomeStore := storapp.NewIncome(db)

// incomeSvc := servicapp.NewIncome(incomeStore)

// incH := handler.NewIncome(incomeSvc)

// app.POST("/users/{user_id}/incomes", incH.Create)
// app.GET("/users/{user_id}/incomes/:id", incH.Get)
// app.PUT("/users/{user_id}/incomes/:id", incH.Update)
// app.DELETE("/users/{user_id}/incomes/:id", incH.Delete)

// fs := http.FileServer(http.Dir("./web/assets/"))
// r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

// r.HandleFunc("/", func(ctx echo.Context) {
// 	t, _ := templatapp.ParseFiles("./web/pages-login.html")
// 	t.Execute(w, nil)
// })

// r.HandleFunc("/register", func(ctx echo.Context) {
// 	t, _ := templatapp.ParseFiles("./web/pages-register.html")
// 	t.Execute(w, nil)
// })
// }
