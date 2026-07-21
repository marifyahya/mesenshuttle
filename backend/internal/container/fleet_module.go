package container

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mesenshuttle-backend/internal/controllers"
	"mesenshuttle-backend/internal/repositories"
	"mesenshuttle-backend/internal/services"
)

type FleetModule struct {
	controller *controllers.FleetController
}

func NewFleetModule(db *gorm.DB) *FleetModule {
	repo := repositories.NewFleetRepository(db)
	svc := services.NewFleetService(repo)
	ctrl := controllers.NewFleetController(svc)

	return &FleetModule{controller: ctrl}
}

func (m *FleetModule) RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	fleetGroup := protected.Group("/fleets")
	fleetGroup.GET("", m.controller.GetFleets)
	fleetGroup.POST("", m.controller.CreateFleet)
	fleetGroup.PUT("/:id", m.controller.UpdateFleet)
}
