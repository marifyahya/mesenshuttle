package services

import (
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/repositories"
	"mesenshuttle-backend/pkg/apperrors"
)

type FleetService interface {
	GetAllFleets(page, limit int) ([]models.Fleet, int64, error)
	CreateFleet(fleet *models.Fleet) error
}

type fleetService struct {
	fleetRepo repositories.FleetRepository
}

func NewFleetService(fleetRepo repositories.FleetRepository) FleetService {
	return &fleetService{fleetRepo}
}

func (s *fleetService) GetAllFleets(page, limit int) ([]models.Fleet, int64, error) {
	return s.fleetRepo.FindAll(page, limit)
}

func (s *fleetService) CreateFleet(fleet *models.Fleet) error {
	if s.fleetRepo.ExistsByPlateNumber(fleet.PlateNumber) {
		return apperrors.ErrDuplicatePlateNumber
	}
	return s.fleetRepo.Create(fleet)
}
