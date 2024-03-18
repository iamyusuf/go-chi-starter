package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type ChiServer struct {
	mux *chi.Mux
}

func (s ChiServer) registerMiddlewares() {
	s.mux.Use(middleware.Logger)
	s.mux.Use(middleware.Recoverer)
	s.mux.Use(middleware.RequestID)
}

func (s ChiServer) registerRoutes() {
	s.mux.Method("GET", "/", Handler(textHandler))
	s.mux.Method("GET", "/custom", Handler(customHandler))
	s.mux.Method("GET", "/test", Handler(testHandler))
}

func (s ChiServer) Start(address string) error {
	s.registerMiddlewares()
	s.registerRoutes()
	return http.ListenAndServe(address, s.mux)
}
