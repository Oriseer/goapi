package api

import (
	"encoding/json"
	"net/http"
)

// Money Balance Params
type MoneyBalanceParams struct {
	Username string
}

// Money balance Response
type MoneyBalanceResponse struct {
	// Account Balance
	Balance int64

	// Error Message Code
	Code int
}

// Error Response
type ErrorResponse struct {
	// Error Message
	Message string

	// Error Message Code
	code int
}

func writeError(w http.ResponseWriter, message string, code int) {
	var response = ErrorResponse{
		Message: message,
		code:    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

var (
	RequstErrorHandler = func(w http.ResponseWriter, error error) {
		writeError(w, error.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal Error Occured", http.StatusInternalServerError)
	}
)
