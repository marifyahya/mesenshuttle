package utils

import (
	"github.com/gin-gonic/gin"
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
