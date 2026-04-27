package todo

import (
	"encoding/json"
	"net/http"
	"errors"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, message string, status int) {
	errorMessage := ErrorMessage{
		Error: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorMessage)
}

func WriteResponse(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func HandleServiceError(err error) {
	var w http.ResponseWriter
	var se *ServiceError
	if errors.As(err, &se) {
		switch se.Type {
		case VALIDATION:
			WriteError(w, se.Message, http.StatusBadRequest)
			return
		case CONFLICT:
			WriteError(w, se.Message, http.StatusConflict)
			return
		default:
			WriteError(w, se.Message, http.StatusInternalServerError)
			return
		}
	} else {
		WriteError(w, "une erreur est survenue", http.StatusInternalServerError)
		return
	}
}
