package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type ChiServer struct {
	mux *chi.Mux
}

func (server ChiServer) register(method, pattern string, handler Handler) {
	server.mux.Method(method, pattern, handler)
}

func (server ChiServer) registerMiddlewares() {
	server.mux.Use(middleware.Logger)
	server.mux.Use(middleware.Recoverer)
	server.mux.Use(middleware.RequestID)
}

func (server ChiServer) registerRoutes() {
	server.register("get", "/", server.textHandler)
	server.register("get", "/custom", server.customHandler)
	server.register("get", "/test", server.testHandler)
}

func (server ChiServer) Start(address string) error {
	server.registerMiddlewares()
	server.registerRoutes()
	return http.ListenAndServe(address, server.mux)
}
