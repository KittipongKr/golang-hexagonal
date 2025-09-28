package errs

import (
	"errors"
	"net/http"
)

type App struct {
	Error      error
	Message    string
	StatusCode int
}

var (
	// 5xx - Server Errors
	ErrInternalServer     = App{errors.New("internal server error"), "Internal Server Error", http.StatusInternalServerError}
	ErrServiceUnavailable = App{errors.New("service unavailable"), "Service Unavailable", http.StatusServiceUnavailable}
	ErrDatabase           = App{errors.New("database error"), "Database Error", http.StatusInternalServerError}

	// 4xx - Client Errors
	ErrBadRequest   = App{errors.New("bad request"), "Bad Request", http.StatusBadRequest}
	ErrUnauthorized = App{errors.New("unauthorized"), "Unauthorized", http.StatusUnauthorized}
	ErrForbidden    = App{errors.New("forbidden"), "Forbidden", http.StatusForbidden}
	ErrNotFound     = App{errors.New("data not found"), "Data Not Found", http.StatusNotFound}
	ErrConflict     = App{errors.New("conflict"), "Conflict", http.StatusConflict}

	// Validation
	ErrInvalidInput     = App{errors.New("invalid input"), "Invalid Input", http.StatusBadRequest}
	ErrMissingQuery     = App{errors.New("missing required query string"), "Missing Required Query string", http.StatusBadRequest}
	ErrInvalidQuery     = App{errors.New("invalid query string"), "Invalid Query string", http.StatusBadRequest}
	ErrMissingParams    = App{errors.New("missing required param string"), "Missing Required Param string", http.StatusBadRequest}
	ErrValidationFailed = App{errors.New("validation failed"), "validation failed", http.StatusUnprocessableEntity}

	// Auth
	ErrTokenExpired = App{errors.New("token expired"), "Token Expired", http.StatusUnauthorized}
	ErrInvalidToken = App{errors.New("invalid expired"), "Invalid Token", http.StatusUnauthorized}

	// File
	ErrFileTooLarge = App{errors.New("file Too Large"), "File Too Large", http.StatusRequestEntityTooLarge}
	ErrUploadFailed = App{errors.New("file upload failed"), "File Upload Failed", http.StatusInternalServerError}
)
