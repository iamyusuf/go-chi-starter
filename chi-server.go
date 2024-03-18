package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type App struct {
	mux *chi.Mux
}

func (app App) register(method, pattern string, handler Handler) {
	app.mux.Method(method, pattern, handler)
}

func (app App) registerMiddlewares() {
	app.mux.Use(middleware.Logger)
	app.mux.Use(middleware.Recoverer)
	app.mux.Use(middleware.RequestID)
}

func (app App) registerRoutes() {
	app.register("get", "/", app.textHandler)
	app.register("get", "/custom", app.customHandler)
	app.register("get", "/test", app.testHandler)
}

func (app App) Start(address string) error {
	app.registerMiddlewares()
	app.registerRoutes()
	return http.ListenAndServe(address, app.mux)
}
