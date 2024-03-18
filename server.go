package main

import "github.com/go-chi/chi/v5"

type Server interface {
	Start(address string) error
}

func NewServer() Server {
	s := ChiServer{
		mux: chi.NewRouter(),
	}

	return s
}
