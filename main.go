package main

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		_, err := w.Write([]byte("bad"))
		if err != nil {
			return
		}
	}
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}

	_, err := w.Write([]byte("foo"))
	return err
}

type ChiServer struct {
	mux *chi.Mux
}

func (s ChiServer) Start(address string) error {
	return http.ListenAndServe(address, s.mux)
}

func (s ChiServer) registerMiddlewares() {
	s.mux.Use(middleware.Logger)
	s.mux.Use(middleware.Recoverer)
	s.mux.Use(middleware.RequestID)
}

func (s ChiServer) registerRoutes() {
	s.mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		textResponse(w, "Hello World", 200)
	})

	s.mux.Method("GET", "/custom", Handler(customHandler))
}

type Server interface {
	Start(address string) error
}

func NewServer() Server {
	s := ChiServer{
		mux: chi.NewRouter(),
	}

	return s
}

func main() {
	s := NewServer()
	err := s.Start(":3000")

	if err != nil {
		return
	}
}

func textResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(message))

	if err != nil {
		return
	}
}
