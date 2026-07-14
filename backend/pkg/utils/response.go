package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mesenshuttle-backend/internal/dto"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, errorMsg string) {
	c.JSON(statusCode, APIResponse{
		Status: "error",
		Error:  errorMsg,
	})
}

func ValidationErrorResponse(c *gin.Context, statusCode int, errors interface{}) {
	c.JSON(statusCode, APIResponse{
		Status: "error",
		Error:  "Validation failed",
		Data:   errors,
	})
}

func FormatValidationErrors(err error) []dto.FieldError {
	var fieldErrs []dto.FieldError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			fieldErrs = append(fieldErrs, dto.FieldError{Field: fe.Field(), Message: fe.Tag()})
		}
	} else {
		fieldErrs = append(fieldErrs, dto.FieldError{Field: "payload", Message: err.Error()})
	}
	return fieldErrs
}
