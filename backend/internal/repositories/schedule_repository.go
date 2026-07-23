package repositories

import (
	"mesenshuttle-backend/internal/models"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	FindAll(page, limit int) ([]models.Schedule, int64, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db}
}

func (r *scheduleRepository) FindAll(page, limit int) ([]models.Schedule, int64, error) {
	var schedules []models.Schedule
	var totalCount int64

	if err := r.db.Model(&models.Schedule{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Preload("Route").Preload("Fleet").Limit(limit).Offset(offset).Find(&schedules).Error; err != nil {
		return nil, 0, err
	}
	return schedules, totalCount, nil
}
