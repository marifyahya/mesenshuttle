package routes

import (
	"mesenshuttle-backend/internal/controllers"
	"mesenshuttle-backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authController *controllers.AuthController, routeController *controllers.RouteController, jwtSecret string) *gin.Engine {
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

			protected := adminGroup.Group("/")
			protected.Use(middlewares.JWTAuthMiddleware(jwtSecret))
			{
				protected.GET("/routes", routeController.GetRoutes)
				protected.POST("/routes", routeController.CreateRoute)
			}
		}
	}

	return r
}
