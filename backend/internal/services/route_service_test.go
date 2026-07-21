package services_test

import (
	"errors"
	"testing"

	"mesenshuttle-backend/internal/dto"
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

func (m *MockRouteRepository) FindByID(id string) (*models.Route, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Route), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRouteRepository) Update(route *models.Route) error {
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

func TestRouteService_UpdateRoute(t *testing.T) {
	mockRepo := new(MockRouteRepository)
	routeService := services.NewRouteService(mockRepo)
	
	routeID := uuid.New().String()

	t.Run("Success", func(t *testing.T) {
		existingRoute := &models.Route{OriginCity: "Jakarta", DestinationCity: "Bandung"}
		input := &dto.UpdateRouteRequest{
			OriginCity:      "New Jakarta",
			OriginPool:      "New Pool JKT",
			DestinationCity: "New Bandung",
			DestinationPool: "New Pool BDG",
		}
		
		mockRepo.On("FindByID", routeID).Return(existingRoute, nil).Once()
		mockRepo.On("Update", existingRoute).Return(nil).Once()

		route, err := routeService.UpdateRoute(routeID, input)

		assert.NoError(t, err)
		assert.Equal(t, "New Jakarta", route.OriginCity)
		assert.Equal(t, "New Pool BDG", route.DestinationPool)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		input := &dto.UpdateRouteRequest{}
		mockRepo.On("FindByID", routeID).Return(nil, errors.New("record not found")).Once()

		route, err := routeService.UpdateRoute(routeID, input)

		assert.Error(t, err)
		assert.Nil(t, route)
		assert.Equal(t, "Route not found", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Update Error", func(t *testing.T) {
		existingRoute := &models.Route{OriginCity: "Jakarta"}
		input := &dto.UpdateRouteRequest{}
		mockRepo.On("FindByID", routeID).Return(existingRoute, nil).Once()
		mockRepo.On("Update", existingRoute).Return(errors.New("db update error")).Once()

		route, err := routeService.UpdateRoute(routeID, input)

		assert.Error(t, err)
		assert.Nil(t, route)
		assert.Equal(t, "db update error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func (m *MockRouteRepository) Delete(route *models.Route) error {
	args := m.Called(route)
	return args.Error(0)
}

func TestRouteService_DeleteRoute(t *testing.T) {
	mockRepo := new(MockRouteRepository)
	routeService := services.NewRouteService(mockRepo)
	
	routeID := uuid.New().String()

	t.Run("Success", func(t *testing.T) {
		existingRoute := &models.Route{ID: uuid.MustParse(routeID)}
		
		mockRepo.On("FindByID", routeID).Return(existingRoute, nil).Once()
		mockRepo.On("Delete", existingRoute).Return(nil).Once()

		err := routeService.DeleteRoute(routeID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		mockRepo.On("FindByID", routeID).Return(nil, errors.New("record not found")).Once()

		err := routeService.DeleteRoute(routeID)

		assert.Error(t, err)
		assert.Equal(t, "Route not found", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Delete Error", func(t *testing.T) {
		existingRoute := &models.Route{ID: uuid.MustParse(routeID)}
		mockRepo.On("FindByID", routeID).Return(existingRoute, nil).Once()
		mockRepo.On("Delete", existingRoute).Return(errors.New("db delete error")).Once()

		err := routeService.DeleteRoute(routeID)

		assert.Error(t, err)
		assert.Equal(t, "db delete error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
