package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kholeur9/dhaclub-app/internal/apperrors"
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

func HandleServiceError(w http.ResponseWriter, err error) {
	var se *apperrors.ServiceError
	if errors.As(err, &se) {
		switch se.Type {
		case apperrors.VALIDATION:
			WriteError(w, se.Message, http.StatusBadRequest)
			return
		case apperrors.CONFLICT:
			WriteError(w, se.Message, http.StatusConflict)
			return
		case apperrors.NOT_FOUND:
			WriteError(w, se.Message, http.StatusConflict)
			return
		default:
			WriteError(w, se.Message, http.StatusInternalServerError)
			return
		}
	} else {
		WriteError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
