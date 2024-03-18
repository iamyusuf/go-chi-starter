package main

import "net/http"

func textResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(message))

	if err != nil {
		return
	}
}
