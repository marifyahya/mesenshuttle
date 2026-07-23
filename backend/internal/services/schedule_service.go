package services

import (
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/repositories"
)

type ScheduleService interface {
	GetAllSchedules(page, limit int) ([]models.Schedule, int64, error)
}

type scheduleService struct {
	scheduleRepo repositories.ScheduleRepository
}

func NewScheduleService(scheduleRepo repositories.ScheduleRepository) ScheduleService {
	return &scheduleService{scheduleRepo}
}

func (s *scheduleService) GetAllSchedules(page, limit int) ([]models.Schedule, int64, error) {
	return s.scheduleRepo.FindAll(page, limit)
}
