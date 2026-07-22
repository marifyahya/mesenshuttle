package repositories

import (
	"mesenshuttle-backend/internal/models"

	"gorm.io/gorm"
)

type FleetRepository interface {
	FindAll(page, limit int) ([]models.Fleet, int64, error)
	FindByID(id string) (*models.Fleet, error)
	ExistsByPlateNumber(plateNumber string) bool
	ExistsByPlateNumberExcludingID(plateNumber, excludeID string) bool
	Create(fleet *models.Fleet) error
	Update(fleet *models.Fleet) error
	Delete(fleet *models.Fleet) error
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

func (r *fleetRepository) FindByID(id string) (*models.Fleet, error) {
	var fleet models.Fleet
	if err := r.db.First(&fleet, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &fleet, nil
}

func (r *fleetRepository) ExistsByPlateNumberExcludingID(plateNumber, excludeID string) bool {
	var count int64
	r.db.Model(&models.Fleet{}).Where("plate_number = ? AND id != ?", plateNumber, excludeID).Count(&count)
	return count > 0
}

func (r *fleetRepository) Update(fleet *models.Fleet) error {
	if err := r.db.Save(fleet).Error; err != nil {
		return err
	}
	return nil
}

func (r *fleetRepository) Delete(fleet *models.Fleet) error {
	if err := r.db.Delete(fleet).Error; err != nil {
		return err
	}
	return nil
}
