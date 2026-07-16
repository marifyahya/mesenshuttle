package repositories

import (
	"mesenshuttle-backend/internal/models"

	"gorm.io/gorm"
)

type FleetRepository interface {
	FindAll(page, limit int) ([]models.Fleet, int64, error)
	ExistsByPlateNumber(plateNumber string) bool
	Create(fleet *models.Fleet) error
}

type fleetRepository struct {
	db *gorm.DB
}

func NewFleetRepository(db *gorm.DB) FleetRepository {
	return &fleetRepository{db}
}

func (r *fleetRepository) FindAll(page, limit int) ([]models.Fleet, int64, error) {
	var fleets []models.Fleet
	var totalCount int64

	if err := r.db.Model(&models.Fleet{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Limit(limit).Offset(offset).Find(&fleets).Error; err != nil {
		return nil, 0, err
	}
	return fleets, totalCount, nil
}

func (r *fleetRepository) ExistsByPlateNumber(plateNumber string) bool {
	var count int64
	r.db.Model(&models.Fleet{}).Where("plate_number = ?", plateNumber).Count(&count)
	return count > 0
}

func (r *fleetRepository) Create(fleet *models.Fleet) error {
	if err := r.db.Create(fleet).Error; err != nil {
		return err
	}
	return nil
}
