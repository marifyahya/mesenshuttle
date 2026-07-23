package container

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mesenshuttle-backend/internal/controllers"
	"mesenshuttle-backend/internal/repositories"
	"mesenshuttle-backend/internal/services"
)

type ScheduleModule struct {
	controller *controllers.ScheduleController
}

func NewScheduleModule(db *gorm.DB) *ScheduleModule {
	repo := repositories.NewScheduleRepository(db)
	svc := services.NewScheduleService(repo)
	ctrl := controllers.NewScheduleController(svc)

	return &ScheduleModule{controller: ctrl}
}

func (m *ScheduleModule) RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup) {
	scheduleGroup := protected.Group("/schedules")
	scheduleGroup.GET("", m.controller.GetSchedules)
}
