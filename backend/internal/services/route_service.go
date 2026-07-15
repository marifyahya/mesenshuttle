package services

import (
	"mesenshuttle-backend/internal/dto"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/repositories"
)

type RouteService interface {
	GetAllRoutes(page, limit int) ([]models.Route, int64, error)
	CreateRoute(route *models.Route) error
	UpdateRoute(id string, input *dto.UpdateRouteRequest) (*models.Route, error)
	DeleteRoute(id string) error
}

type routeService struct {
	routeRepo repositories.RouteRepository
}

func NewRouteService(routeRepo repositories.RouteRepository) RouteService {
	return &routeService{routeRepo}
}

func (s *routeService) GetAllRoutes(page, limit int) ([]models.Route, int64, error) {
	return s.routeRepo.FindAll(page, limit)
}

func (s *routeService) CreateRoute(route *models.Route) error {
	return s.routeRepo.Create(route)
}

func (s *routeService) UpdateRoute(id string, input *dto.UpdateRouteRequest) (*models.Route, error) {
	route, err := s.routeRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	route.OriginCity = input.OriginCity
	route.OriginPool = input.OriginPool
	route.DestinationCity = input.DestinationCity
	route.DestinationPool = input.DestinationPool

	if err := s.routeRepo.Update(route); err != nil {
		return nil, err
	}

	return route, nil
}

func (s *routeService) DeleteRoute(id string) error {
	route, err := s.routeRepo.FindByID(id)
	if err != nil {
		return err
	}

	if err := s.routeRepo.Delete(route); err != nil {
		return err
	}

	return nil
}
