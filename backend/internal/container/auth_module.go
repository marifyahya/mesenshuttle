package container

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	
	"mesenshuttle-backend/internal/config"
	"mesenshuttle-backend/internal/controllers"
	"mesenshuttle-backend/internal/middlewares"
	"mesenshuttle-backend/internal/repositories"
	"mesenshuttle-backend/internal/services"
)

type AuthModule struct {
	controller *controllers.AuthController
	redis      *redis.Client
}

func NewAuthModule(db *gorm.DB, cfg *config.Config, redisClient *redis.Client) *AuthModule {
	repo := repositories.NewAdminRepository(db)
	svc := services.NewAuthService(repo, cfg)
	ctrl := controllers.NewAuthController(svc)
	
	return &AuthModule{controller: ctrl, redis: redisClient}
}

func (m *AuthModule) RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	public.POST("/login", middlewares.RateLimiter(m.redis, 5, time.Minute), m.controller.Login)
}
