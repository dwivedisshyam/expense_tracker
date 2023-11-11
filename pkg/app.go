package pkg

import (
	"log"
	"net/http"

	"github.com/dwivedisshyam/go-lib/pkg/config"
	"github.com/gorilla/mux"
)

type App struct {
	Config config.Config
	router *mux.Router
}

func NewApp() *App {
	return &App{
		Config: config.New(),
		router: mux.NewRouter().StrictSlash(false),
	}
}

func (app *App) GetEnv(key string) string {
	return app.Config.Get(key)
}

func (app *App) addRoute(method, pattern string, handler http.Handler) {
	app.router.NewRoute().Methods(method).Path(pattern).Handler(handler)
}

func (app *App) GET(pattern string, handler http.HandlerFunc) {
	app.addRoute(http.MethodGet, pattern, handler)
}

func (app *App) POST(pattern string, handler http.HandlerFunc) {
	app.addRoute(http.MethodPost, pattern, handler)
}

func (app *App) PUT(pattern string, handler http.HandlerFunc) {
	app.addRoute(http.MethodPut, pattern, handler)
}

func (app *App) DELETE(pattern string, handler http.HandlerFunc) {
	app.addRoute(http.MethodDelete, pattern, handler)
}

func (app *App) Use(mwf ...mux.MiddlewareFunc) {
	app.router.Use(mwf...)
}

func (app *App) Run() {
	log.Println("Starting HTTP server @ :8000")
	log.Fatal(http.ListenAndServe(":8000", app.router))
}
