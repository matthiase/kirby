package httputil

import (
	"encoding/json"
	"kirby/errors"
	"net/http"
)

// SuccessResponse envelope
type SuccessResponse struct {
	Data interface{} `json:"data,omitempty" `
}

// RespondWithStatus responds with status code
func RespondWithStatus(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

// RespondWithJSON responds with seralized JSON response
func RespondWithJSON(w http.ResponseWriter, httpStatus int, data interface{}) {
	response := SuccessResponse{data}
	json, err := json.Marshal(response)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(json)
}

// RespondWithError responds with error status code and serialized JSON error message
func RespondWithError(w http.ResponseWriter, httpStatus uint32, err error) {
	errors := []errors.ApplicationError{errors.ApplicationError{Message: err.Error()}}
	json, _ := json.Marshal(errors)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(httpStatus))
	w.Write(json)
}
