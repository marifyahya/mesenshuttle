package utils

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"mesenshuttle-backend/pkg/apperrors"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

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

func ValidationErrorResponse(c *gin.Context, statusCode int, details interface{}) {
	c.JSON(statusCode, APIResponse{
		Status: "error",
		Error:  "Validation failed",
		Data:   details,
	})
}

func HandleError(c *gin.Context, err error) {
	var appErr *apperrors.AppError

	if errors.As(err, &appErr) {
		if len(appErr.ValidationDetails) > 0 {
			ValidationErrorResponse(c, appErr.Code, appErr.ValidationDetails)
		} else {
			ErrorResponse(c, appErr.Code, appErr.Message)
		}
		return
	}

	ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	case "oneof":
		return fmt.Sprintf("must be one of: %s", strings.ReplaceAll(fe.Param(), " ", ", "))
	case "min":
		return fmt.Sprintf("must be at least %s", fe.Param())
	case "max":
		return fmt.Sprintf("must be at most %s", fe.Param())
	case "email":
		return "must be a valid email address"
	default:
		return fe.Tag()
	}
}

func FormatValidationErrors(err error) []apperrors.ValidationDetail {
	var fieldErrs []apperrors.ValidationDetail
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			fieldErrs = append(fieldErrs, apperrors.ValidationDetail{Field: fe.Field(), Message: getErrorMsg(fe)})
		}
	} else {
		fieldErrs = append(fieldErrs, apperrors.ValidationDetail{Field: "payload", Message: err.Error()})
	}
	return fieldErrs
}
