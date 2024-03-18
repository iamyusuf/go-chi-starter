package main

import (
	"errors"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := ""

	if err := h(w, r); err != nil {
		switch {
		case errors.Is(err, ErrInternalServerError):
			// Handle internal server error
			// ...
		case errors.Is(err, ErrBadRequest):
			// Handle bad request
			// ...
		case errors.Is(err, ErrUnauthorized):
			// Handle unauthorized access
			// ...
		case errors.Is(err, ErrForbidden):
			// Handle forbidden access
			// ...
		case errors.Is(err, ErrNotFound):
			// Handle not found error
			// ...
		default:
			// Handle other errors
			// ...
		}

		response := map[string]interface{}{
			"success": true,
			"error":   err,
			"message": message,
		}

		err := jsonResponse(w, 404, response)

		if err != nil {
			return
		}
	}
}

func (server ChiServer) customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}

	_, err := w.Write([]byte("foo"))
	return err
}

func (server ChiServer) testHandler(w http.ResponseWriter, r *http.Request) error {
	responseData := map[string]interface{}{
		"success": true,
		"message": "Successful",
	}

	return jsonResponse(w, http.StatusOK, responseData)
}

func (server ChiServer) textHandler(w http.ResponseWriter, r *http.Request) error {
	return textResponse(w, http.StatusOK, "Hello World")
}
