package repositories

import (
	"mesenshuttle-backend/internal/models"

	"gorm.io/gorm"
)

type RouteRepository interface {
	FindAll(page, limit int) ([]models.Route, int64, error)
	Create(route *models.Route) error
	FindByID(id string) (*models.Route, error)
	Update(route *models.Route) error
}

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) RouteRepository {
	return &routeRepository{db}
}

func (r *routeRepository) FindAll(page, limit int) ([]models.Route, int64, error) {
	var routes []models.Route
	var totalCount int64

	if err := r.db.Model(&models.Route{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := r.db.Limit(limit).Offset(offset).Find(&routes).Error; err != nil {
		return nil, 0, err
	}
	return routes, totalCount, nil
}

func (r *routeRepository) Create(route *models.Route) error {
	if err := r.db.Create(route).Error; err != nil {
		return err
	}
	return nil
}

func (r *routeRepository) FindByID(id string) (*models.Route, error) {
	var route models.Route
	if err := r.db.Where("id = ?", id).First(&route).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func (r *routeRepository) Update(route *models.Route) error {
	if err := r.db.Save(route).Error; err != nil {
		return err
	}
	return nil
}
