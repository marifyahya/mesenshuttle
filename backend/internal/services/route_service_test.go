package services_test

import (
	"errors"
	"testing"

	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/services"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRouteRepository struct {
	mock.Mock
}

func (m *MockRouteRepository) FindAll(page, limit int) ([]models.Route, int64, error) {
	args := m.Called(page, limit)
	if args.Get(0) != nil {
		return args.Get(0).([]models.Route), args.Get(1).(int64), args.Error(2)
	}
	return nil, 0, args.Error(2)
}

func (m *MockRouteRepository) Create(route *models.Route) error {
	args := m.Called(route)
	return args.Error(0)
}

func TestRouteService_GetAllRoutes(t *testing.T) {
	mockRepo := new(MockRouteRepository)
	routeService := services.NewRouteService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockRoutes := []models.Route{
			{ID: uuid.New(), OriginCity: "Jakarta", DestinationCity: "Bandung"},
		}
		mockRepo.On("FindAll", 1, 10).Return(mockRoutes, int64(1), nil).Once()

		routes, total, err := routeService.GetAllRoutes(1, 10)

		assert.NoError(t, err)
		assert.NotNil(t, routes)
		assert.Equal(t, int64(1), total)
		assert.Len(t, routes, 1)
		assert.Equal(t, "Jakarta", routes[0].OriginCity)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		mockRepo.On("FindAll", 1, 10).Return(nil, int64(0), errors.New("db error")).Once()

		routes, total, err := routeService.GetAllRoutes(1, 10)

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		assert.Nil(t, routes)
		assert.Equal(t, int64(0), total)
		mockRepo.AssertExpectations(t)
	})
}

func TestRouteService_CreateRoute(t *testing.T) {
	mockRepo := new(MockRouteRepository)
	routeService := services.NewRouteService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		route := &models.Route{OriginCity: "Jakarta", DestinationCity: "Bandung"}
		mockRepo.On("Create", route).Return(nil).Once()

		err := routeService.CreateRoute(route)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		route := &models.Route{OriginCity: "Jakarta", DestinationCity: "Bandung"}
		mockRepo.On("Create", route).Return(errors.New("db error")).Once()

		err := routeService.CreateRoute(route)

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
