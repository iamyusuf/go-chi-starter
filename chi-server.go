package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strings"
)

type ChiServer struct {
	mux *chi.Mux
}

func (server ChiServer) method(method, pattern string, handler Handler) {
	server.mux.Method(strings.ToUpper(method), pattern, handler)
}

func (server ChiServer) registerMiddlewares() {
	server.mux.Use(middleware.Logger)
	server.mux.Use(middleware.Recoverer)
	server.mux.Use(middleware.RequestID)
}

func (server ChiServer) registerRoutes() {
	server.method("get", "/", server.textHandler)
	server.method("get", "/custom", server.customHandler)
	server.method("get", "/test", server.testHandler)
}

func (server ChiServer) Start(address string) error {
	server.registerMiddlewares()
	server.registerRoutes()
	return http.ListenAndServe(address, server.mux)
}
