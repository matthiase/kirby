package httputil

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse envelope
type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data,omitempty" `
}

// ErrorResponse envelope
type ErrorResponse struct {
	Success bool      `json:"success" example:"false"`
	Error   HTTPError `json:"error,omitempty"`
}

// HTTPError struct
type HTTPError struct {
	Code    uint32 `json:"code" example:"40001"`
	Message string `json:"message" example:"status bad request"`
}

// RespondWithStatus responds with status code
func RespondWithStatus(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}

// RespondWithJSON responds with seralized JSON response
func RespondWithJSON(w http.ResponseWriter, httpStatus int, data interface{}) {
	response := SuccessResponse{
		Success: true,
		Data:    data,
	}
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
	resp := ErrorResponse{
		Success: false,
		Error: HTTPError{
			Code:    httpStatus,
			Message: err.Error(),
		},
	}
	json, err := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(httpStatus))
	w.Write(json)
}
