package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ChiServer struct {
	mux *chi.Mux
}

func (s ChiServer) Start(address string) error {
	return http.ListenAndServe(address, s.mux)
}
