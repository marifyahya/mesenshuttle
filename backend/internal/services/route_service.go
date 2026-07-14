package services

import (
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/repositories"
)

type RouteService interface {
	GetAllRoutes(page, limit int) ([]models.Route, int64, error)
	CreateRoute(route *models.Route) error
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
