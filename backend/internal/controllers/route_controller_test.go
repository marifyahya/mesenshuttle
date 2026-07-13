package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"mesenshuttle-backend/internal/controllers"
	"mesenshuttle-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRouteService struct {
	mock.Mock
}

func (m *MockRouteService) GetAllRoutes() ([]models.Route, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]models.Route), args.Error(1)
	}
	return nil, args.Error(1)
}

func setupRouteRouter(routeService *MockRouteService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	routeController := controllers.NewRouteController(routeService)
	r.GET("/api/admin/routes", routeController.GetRoutes)
	return r
}

func TestRouteController_GetRoutes(t *testing.T) {
	mockService := new(MockRouteService)
	r := setupRouteRouter(mockService)

	t.Run("Success", func(t *testing.T) {
		mockRoutes := []models.Route{
			{ID: uuid.New(), OriginCity: "Jakarta", DestinationCity: "Bandung"},
		}
		mockService.On("GetAllRoutes").Return(mockRoutes, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/admin/routes", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Routes retrieved successfully", response["message"])
		
		data := response["data"].([]interface{})
		assert.Len(t, data, 1)
		
		route := data[0].(map[string]interface{})
		assert.Equal(t, "Jakarta", route["OriginCity"])

		mockService.AssertExpectations(t)
	})

	t.Run("Service Error", func(t *testing.T) {
		mockService.On("GetAllRoutes").Return(nil, errors.New("service error")).Once()

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
