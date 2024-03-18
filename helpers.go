package main

import (
	"encoding/json"
	"net/http"
)

func textResponse(w http.ResponseWriter, status int, message string) error {
	w.WriteHeader(status)
	_, err := w.Write([]byte(message))
	return err
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return err
	}

	return nil
}
