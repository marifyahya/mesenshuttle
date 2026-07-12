package routes

import (
	"mesenshuttle-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authController *controllers.AuthController) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Mesenshuttle Backend is running!",
		})
	})

	api := r.Group("/api")
	{
		adminGroup := api.Group("/admin")
		{
			adminGroup.POST("/login", authController.Login)
		}
	}

	return r
}
