package container

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	
	"mesenshuttle-backend/internal/controllers"
	"mesenshuttle-backend/internal/repositories"
	"mesenshuttle-backend/internal/services"
)

type RouteModule struct {
	controller *controllers.RouteController
}

func NewRouteModule(db *gorm.DB) *RouteModule {
	repo := repositories.NewRouteRepository(db)
	svc := services.NewRouteService(repo)
	ctrl := controllers.NewRouteController(svc)
	
	return &RouteModule{controller: ctrl}
}

func (m *RouteModule) RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	routeGroup := protected.Group("/routes")
	routeGroup.GET("", m.controller.GetRoutes)
	routeGroup.POST("", m.controller.CreateRoute)
	routeGroup.PUT("/:id", m.controller.UpdateRoute)
}
