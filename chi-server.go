package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type ChiServer struct {
	mux *chi.Mux
}

func (server ChiServer) registerMiddlewares() {
	server.mux.Use(middleware.Logger)
	server.mux.Use(middleware.Recoverer)
	server.mux.Use(middleware.RequestID)
}

func (server ChiServer) registerRoutes() {
	server.mux.Method("GET", "/", Handler(server.textHandler))
	server.mux.Method("GET", "/custom", Handler(server.customHandler))
	server.mux.Method("GET", "/test", Handler(server.testHandler))
}

func (server ChiServer) Start(address string) error {
	server.registerMiddlewares()
	server.registerRoutes()
	return http.ListenAndServe(address, server.mux)
}
