package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, message string, status int) {
	errorMessage := ErrorMessage{
		Error: message,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorMessage)
}

func WriteResponse(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
