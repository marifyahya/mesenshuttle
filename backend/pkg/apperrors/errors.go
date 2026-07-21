package apperrors

import (
	"fmt"
	"net/http"
	"strings"
)

// ValidationDetail represents a single field-level validation error.
type ValidationDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// AppError is a structured application error carrying an HTTP status code,
// a human-readable message, and optional field-level validation details.
type AppError struct {
	Code              int
	Message           string
	ValidationDetails []ValidationDetail
}

func (e *AppError) Error() string {
	return e.Message
}

// New creates a generic AppError with an explicit HTTP status code and message.
func New(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// NewInternal creates a 500 error. The message is generic to avoid leaking
// internal details to the client.
func NewInternal() *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: "Internal server error"}
}

// NewNotFound creates a 404 error for a missing resource.
func NewNotFound(resource string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("%s not found", resource),
	}
}

// NewDuplicateField creates a 400 validation error for a duplicate unique field.
func NewDuplicateField(field string) *AppError {
	humanReadableField := strings.ReplaceAll(field, "_", " ")
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
		ValidationDetails: []ValidationDetail{
			{
				Field:   field,
				Message: fmt.Sprintf("%s is already registered", humanReadableField),
			},
		},
	}
}

// NewUnauthorized creates a 401 error with a custom message.
func NewUnauthorized(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}
