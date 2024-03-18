package main

import (
	"errors"
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

func main() {
	s := NewServer()
	err := s.Start(":3000")

	if err != nil {
		return
	}
}
