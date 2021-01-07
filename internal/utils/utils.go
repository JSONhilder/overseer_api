package utils

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Message string `json:"message"`
}

// JSONError - Handle errors with a json response
func JSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	jsonErr := httpError{Message: err}
	json.NewEncoder(w).Encode(jsonErr)
}
