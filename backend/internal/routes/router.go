package routes

import (
	"mesenshuttle-backend/internal/container"
	"mesenshuttle-backend/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetupRouter(modules []container.Module, jwtSecret string, redisClient *redis.Client) *gin.Engine {
	r := gin.New()

	r.Use(middlewares.StructuredLogger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Mesenshuttle Backend is running!",
		})
	})

	api := r.Group("/api")

	adminPublic := api.Group("/admin")
	adminProtected := api.Group("/admin")
	adminProtected.Use(middlewares.JWTAuthMiddleware(jwtSecret))

	for _, m := range modules {
		m.RegisterRoutes(adminPublic, adminProtected)
	}

	return r
}
