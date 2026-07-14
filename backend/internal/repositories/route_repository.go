package repositories

import (
	"mesenshuttle-backend/internal/models"

	"gorm.io/gorm"
)

type RouteRepository interface {
	FindAll() ([]models.Route, error)
	Create(route *models.Route) error
}

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) RouteRepository {
	return &routeRepository{db}
}

func (r *routeRepository) FindAll() ([]models.Route, error) {
	var routes []models.Route
	if err := r.db.Find(&routes).Error; err != nil {
		return nil, err
	}
	return routes, nil
}

func (r *routeRepository) Create(route *models.Route) error {
	if err := r.db.Create(route).Error; err != nil {
		return err
	}
	return nil
}
