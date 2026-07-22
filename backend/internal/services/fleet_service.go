package services

import (
	"mesenshuttle-backend/internal/dto"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/repositories"
	"mesenshuttle-backend/pkg/apperrors"
)

type FleetService interface {
	GetAllFleets(page, limit int) ([]models.Fleet, int64, error)
	CreateFleet(fleet *models.Fleet) error
	UpdateFleet(id string, req *dto.UpdateFleetRequest) (*models.Fleet, error)
	DeleteFleet(id string) error
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
		return apperrors.NewDuplicateField("plate_number")
	}
	return s.fleetRepo.Create(fleet)
}

func (s *fleetService) UpdateFleet(id string, req *dto.UpdateFleetRequest) (*models.Fleet, error) {
	fleet, err := s.fleetRepo.FindByID(id)
	if err != nil {
		return nil, apperrors.NewNotFound("Fleet")
	}

	if req.PlateNumber != nil && *req.PlateNumber != fleet.PlateNumber {
		if s.fleetRepo.ExistsByPlateNumberExcludingID(*req.PlateNumber, id) {
			return nil, apperrors.NewDuplicateField("plate_number")
		}
		fleet.PlateNumber = *req.PlateNumber
	}

	if req.Name != nil {
		fleet.Name = *req.Name
	}
	if req.Type != nil {
		fleet.Type = *req.Type
	}
	if req.TotalSeats != nil {
		fleet.TotalSeats = *req.TotalSeats
	}
	if req.ActiveStatus != nil {
		fleet.ActiveStatus = *req.ActiveStatus
	}

	if err := s.fleetRepo.Update(fleet); err != nil {
		return nil, err
	}

	return fleet, nil
}

func (s *fleetService) DeleteFleet(id string) error {
	fleet, err := s.fleetRepo.FindByID(id)
	if err != nil {
		return apperrors.NewNotFound("Fleet")
	}

	if err := s.fleetRepo.Delete(fleet); err != nil {
		return err
	}

	return nil
}
