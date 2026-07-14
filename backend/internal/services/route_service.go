package services

import (
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/repositories"
)

type RouteService interface {
	GetAllRoutes() ([]models.Route, error)
	CreateRoute(route *models.Route) error
}

type routeService struct {
	routeRepo repositories.RouteRepository
}

func NewRouteService(routeRepo repositories.RouteRepository) RouteService {
	return &routeService{routeRepo}
}

func (s *routeService) GetAllRoutes() ([]models.Route, error) {
	return s.routeRepo.FindAll()
}

func (s *routeService) CreateRoute(route *models.Route) error {
	return s.routeRepo.Create(route)
}
