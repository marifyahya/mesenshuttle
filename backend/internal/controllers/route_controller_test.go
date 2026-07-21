package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"mesenshuttle-backend/internal/controllers"
	"mesenshuttle-backend/internal/dto"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/pkg/apperrors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRouteService struct {
	mock.Mock
}

func (m *MockRouteService) GetAllRoutes(page, limit int) ([]models.Route, int64, error) {
	args := m.Called(page, limit)
	if args.Get(0) != nil {
		return args.Get(0).([]models.Route), args.Get(1).(int64), args.Error(2)
	}
	return nil, 0, args.Error(2)
}

func (m *MockRouteService) CreateRoute(route *models.Route) error {
	args := m.Called(route)
	return args.Error(0)
}

func (m *MockRouteService) UpdateRoute(id string, input *dto.UpdateRouteRequest) (*models.Route, error) {
	args := m.Called(id, input)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Route), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRouteService) DeleteRoute(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupRouteRouter(routeService *MockRouteService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	routeController := controllers.NewRouteController(routeService)
	r.GET("/api/admin/routes", routeController.GetRoutes)
	r.POST("/api/admin/routes", routeController.CreateRoute)
	r.PUT("/api/admin/routes/:id", routeController.UpdateRoute)
	r.DELETE("/api/admin/routes/:id", routeController.DeleteRoute)
	return r
}

func TestRouteController_GetRoutes(t *testing.T) {
	mockService := new(MockRouteService)
	r := setupRouteRouter(mockService)

	t.Run("Success", func(t *testing.T) {
		mockRoutes := []models.Route{
			{ID: uuid.New(), OriginCity: "Jakarta", DestinationCity: "Bandung"},
		}
		mockService.On("GetAllRoutes", 1, 10).Return(mockRoutes, int64(1), nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/admin/routes", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Routes retrieved successfully", response["message"])

		dataMap := response["data"].(map[string]interface{})
		assert.Equal(t, float64(1), dataMap["total_count"])
		assert.Equal(t, float64(1), dataMap["page"])
		assert.Equal(t, float64(10), dataMap["limit"])

		data := dataMap["data"].([]interface{})
		assert.Len(t, data, 1)

		route := data[0].(map[string]interface{})
		assert.Equal(t, "Jakarta", route["origin_city"])

		mockService.AssertExpectations(t)
	})

	t.Run("Service Error", func(t *testing.T) {
		mockService.On("GetAllRoutes", 1, 10).Return(nil, int64(0), errors.New("service error")).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/admin/routes", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Failed to retrieve routes", response["error"])

		mockService.AssertExpectations(t)
	})
}

func TestRouteController_CreateRoute(t *testing.T) {
	mockService := new(MockRouteService)
	r := setupRouteRouter(mockService)

	t.Run("Success", func(t *testing.T) {
		reqBody := `{"origin_city": "Jakarta", "origin_pool": "Pool Kebon Jeruk", "destination_city": "Bandung", "destination_pool": "Pool Pasteur"}`
		mockService.On("CreateRoute", mock.AnythingOfType("*models.Route")).Return(nil).Once()

		req, _ := http.NewRequest(http.MethodPost, "/api/admin/routes", bytes.NewBuffer([]byte(reqBody)))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Route created successfully", response["message"])

		mockService.AssertExpectations(t)
	})

	t.Run("Invalid Payload", func(t *testing.T) {
		reqBody := `{"origin_city": "Jakarta"}`

		req, _ := http.NewRequest(http.MethodPost, "/api/admin/routes", strings.NewReader(reqBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Validation failed", response["error"])
	})

	t.Run("Service Error", func(t *testing.T) {
		reqBody := `{"origin_city": "Jakarta", "origin_pool": "Pool Kebon Jeruk", "destination_city": "Bandung", "destination_pool": "Pool Pasteur"}`
		mockService.On("CreateRoute", mock.AnythingOfType("*models.Route")).Return(errors.New("db error")).Once()

		req, _ := http.NewRequest(http.MethodPost, "/api/admin/routes", bytes.NewBuffer([]byte(reqBody)))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Failed to create route", response["error"])

		mockService.AssertExpectations(t)
	})
}

func TestRouteController_UpdateRoute(t *testing.T) {
	mockService := new(MockRouteService)
	r := setupRouteRouter(mockService)
	routeID := uuid.New().String()

	t.Run("Success", func(t *testing.T) {
		reqBody := `{"origin_city": "New Jakarta", "origin_pool": "New Pool", "destination_city": "New Bandung", "destination_pool": "New Pool 2"}`
		
		expectedRoute := &models.Route{
			ID:              uuid.MustParse(routeID),
			OriginCity:      "New Jakarta",
			OriginPool:      "New Pool",
			DestinationCity: "New Bandung",
			DestinationPool: "New Pool 2",
		}
		
		mockService.On("UpdateRoute", routeID, mock.AnythingOfType("*dto.UpdateRouteRequest")).Return(expectedRoute, nil).Once()

		req, _ := http.NewRequest(http.MethodPut, "/api/admin/routes/"+routeID, bytes.NewBuffer([]byte(reqBody)))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Route updated successfully", response["message"])
		
		data := response["data"].(map[string]interface{})
		assert.Equal(t, "New Jakarta", data["origin_city"])

		mockService.AssertExpectations(t)
	})

	t.Run("Invalid Payload", func(t *testing.T) {
		reqBody := `{"origin_city": "Jakarta"}` // Missing fields

		req, _ := http.NewRequest(http.MethodPut, "/api/admin/routes/"+routeID, strings.NewReader(reqBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Validation failed", response["error"])
	})

	t.Run("Route Not Found", func(t *testing.T) {
		reqBody := `{"origin_city": "Jakarta", "origin_pool": "Pool Kebon Jeruk", "destination_city": "Bandung", "destination_pool": "Pool Pasteur"}`
		mockService.On("UpdateRoute", routeID, mock.AnythingOfType("*dto.UpdateRouteRequest")).Return(nil, apperrors.NewNotFound("Route")).Once()

		req, _ := http.NewRequest(http.MethodPut, "/api/admin/routes/"+routeID, bytes.NewBuffer([]byte(reqBody)))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Route not found", response["error"])

		mockService.AssertExpectations(t)
	})

	t.Run("Service Error", func(t *testing.T) {
		reqBody := `{"origin_city": "Jakarta", "origin_pool": "Pool Kebon Jeruk", "destination_city": "Bandung", "destination_pool": "Pool Pasteur"}`
		mockService.On("UpdateRoute", routeID, mock.AnythingOfType("*dto.UpdateRouteRequest")).Return(nil, errors.New("db error")).Once()

		req, _ := http.NewRequest(http.MethodPut, "/api/admin/routes/"+routeID, bytes.NewBuffer([]byte(reqBody)))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Internal server error", response["error"])

		mockService.AssertExpectations(t)
	})
}

func TestRouteController_DeleteRoute(t *testing.T) {
	mockService := new(MockRouteService)
	r := setupRouteRouter(mockService)
	routeID := uuid.New().String()

	t.Run("Success", func(t *testing.T) {
		mockService.On("DeleteRoute", routeID).Return(nil).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/api/admin/routes/"+routeID, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Route deleted successfully", response["message"])

		mockService.AssertExpectations(t)
	})

	t.Run("Route Not Found", func(t *testing.T) {
		mockService.On("DeleteRoute", routeID).Return(apperrors.NewNotFound("Route")).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/api/admin/routes/"+routeID, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Route not found", response["error"])

		mockService.AssertExpectations(t)
	})

	t.Run("Service Error", func(t *testing.T) {
		mockService.On("DeleteRoute", routeID).Return(errors.New("db error")).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/api/admin/routes/"+routeID, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Internal server error", response["error"])

		mockService.AssertExpectations(t)
	})
}
