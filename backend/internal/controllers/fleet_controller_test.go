package controllers_test

import (
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

type MockFleetService struct {
	mock.Mock
}

func (m *MockFleetService) GetAllFleets(page, limit int) ([]models.Fleet, int64, error) {
	args := m.Called(page, limit)
	if args.Get(0) != nil {
		return args.Get(0).([]models.Fleet), args.Get(1).(int64), args.Error(2)
	}
	return nil, 0, args.Error(2)
}

func (m *MockFleetService) CreateFleet(fleet *models.Fleet) error {
	args := m.Called(fleet)
	return args.Error(0)
}

func (m *MockFleetService) UpdateFleet(id string, req *dto.UpdateFleetRequest) (*models.Fleet, error) {
	args := m.Called(id, req)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Fleet), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockFleetService) DeleteFleet(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupFleetRouter(fleetService *MockFleetService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	fleetController := controllers.NewFleetController(fleetService)
	r.GET("/api/admin/fleets", fleetController.GetFleets)
	r.POST("/api/admin/fleets", fleetController.CreateFleet)
	r.PUT("/api/admin/fleets/:id", fleetController.UpdateFleet)
	r.DELETE("/api/admin/fleets/:id", fleetController.DeleteFleet)
	return r
}

func TestFleetController_GetFleets(t *testing.T) {
	mockService := new(MockFleetService)
	r := setupFleetRouter(mockService)

	t.Run("Success", func(t *testing.T) {
		mockFleets := []models.Fleet{
			{ID: uuid.New(), Name: "Armada 1", PlateNumber: "B 1234 XYZ", Type: "Premium", TotalSeats: 14, ActiveStatus: true},
		}
		mockService.On("GetAllFleets", 1, 10).Return(mockFleets, int64(1), nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/admin/fleets", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Fleets retrieved successfully", response["message"])

		dataMap := response["data"].(map[string]interface{})
		assert.Equal(t, float64(1), dataMap["total_count"])
		assert.Equal(t, float64(1), dataMap["page"])
		assert.Equal(t, float64(10), dataMap["limit"])

		data := dataMap["data"].([]interface{})
		assert.Len(t, data, 1)

		fleet := data[0].(map[string]interface{})
		assert.Equal(t, "Armada 1", fleet["name"])
		assert.Equal(t, "B 1234 XYZ", fleet["plate_number"])
		assert.Equal(t, "Premium", fleet["type"])
		assert.Equal(t, float64(14), fleet["total_seats"])
		assert.Equal(t, true, fleet["active_status"])

		mockService.AssertExpectations(t)
	})

	t.Run("Service Error", func(t *testing.T) {
		mockService.On("GetAllFleets", 1, 10).Return(nil, int64(0), errors.New("service error")).Once()

		req, _ := http.NewRequest(http.MethodGet, "/api/admin/fleets", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Failed to retrieve fleets", response["error"])

		mockService.AssertExpectations(t)
	})
}

func TestFleetController_CreateFleet(t *testing.T) {
	mockService := new(MockFleetService)
	r := setupFleetRouter(mockService)

	t.Run("Success", func(t *testing.T) {
		reqBody := `{"name": "Armada 1", "plate_number": "B 1234 XYZ", "type": "Premium", "total_seats": 14}`
		mockService.On("CreateFleet", mock.AnythingOfType("*models.Fleet")).Return(nil).Once()

		req, _ := http.NewRequest(http.MethodPost, "/api/admin/fleets", strings.NewReader(reqBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Fleet created successfully", response["message"])

		mockService.AssertExpectations(t)
	})

	t.Run("Missing Required Fields", func(t *testing.T) {
		reqBody := `{"name": "Armada 1", "plate_number": "B 1234 XYZ"}`

		req, _ := http.NewRequest(http.MethodPost, "/api/admin/fleets", strings.NewReader(reqBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Validation failed", response["error"])

		data := response["data"].([]interface{})
		assert.Len(t, data, 2)

		field1 := data[0].(map[string]interface{})
		assert.Equal(t, "type", field1["field"])
		assert.Equal(t, "this field is required", field1["message"])

		field2 := data[1].(map[string]interface{})
		assert.Equal(t, "total_seats", field2["field"])
		assert.Equal(t, "this field is required", field2["message"])
	})

	t.Run("Invalid Enum Type", func(t *testing.T) {
		reqBody := `{"name": "Armada 1", "plate_number": "B 1234 XYZ", "type": "Hiace", "total_seats": 14}`

		req, _ := http.NewRequest(http.MethodPost, "/api/admin/fleets", strings.NewReader(reqBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Validation failed", response["error"])

		data := response["data"].([]interface{})
		assert.Len(t, data, 1)

		field1 := data[0].(map[string]interface{})
		assert.Equal(t, "type", field1["field"])
		assert.Equal(t, "must be one of: Reguler, Premium", field1["message"])
	})

	t.Run("Service Error", func(t *testing.T) {
		reqBody := `{"name": "Armada 1", "plate_number": "B 1234 XYZ", "type": "Premium", "total_seats": 14}`
		mockService.On("CreateFleet", mock.AnythingOfType("*models.Fleet")).Return(errors.New("db error")).Once()

		req, _ := http.NewRequest(http.MethodPost, "/api/admin/fleets", strings.NewReader(reqBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Internal server error", response["error"])

		mockService.AssertExpectations(t)
	})

	t.Run("Duplicate Plate Number", func(t *testing.T) {
		reqBody := `{"name": "Armada 2", "plate_number": "B 1234 XYZ", "type": "Premium", "total_seats": 14}`
		mockService.On("CreateFleet", mock.AnythingOfType("*models.Fleet")).Return(apperrors.NewDuplicateField("plate_number")).Once()

		req, _ := http.NewRequest(http.MethodPost, "/api/admin/fleets", strings.NewReader(reqBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Validation failed", response["error"])

		data := response["data"].([]interface{})
		field := data[0].(map[string]interface{})
		assert.Equal(t, "plate_number", field["field"])
		assert.Equal(t, "plate number is already registered", field["message"])

		mockService.AssertExpectations(t)
	})
}

func TestFleetController_UpdateFleet(t *testing.T) {
	mockService := new(MockFleetService)
	r := setupFleetRouter(mockService)

	t.Run("Success", func(t *testing.T) {
		id := uuid.New().String()
		reqBody := `{"name": "Armada 1 Updated", "type": "Premium"}`

		expectedFleet := &models.Fleet{
			ID:   uuid.MustParse(id),
			Name: "Armada 1 Updated",
			Type: "Premium",
		}

		mockService.On("UpdateFleet", id, mock.AnythingOfType("*dto.UpdateFleetRequest")).Return(expectedFleet, nil).Once()

		req, _ := http.NewRequest(http.MethodPut, "/api/admin/fleets/"+id, strings.NewReader(reqBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Fleet updated successfully", response["message"])

		mockService.AssertExpectations(t)
	})

	t.Run("Fleet Not Found", func(t *testing.T) {
		id := uuid.New().String()
		reqBody := `{"name": "Armada 1 Updated"}`

		mockService.On("UpdateFleet", id, mock.AnythingOfType("*dto.UpdateFleetRequest")).Return(nil, apperrors.NewNotFound("Fleet")).Once()

		req, _ := http.NewRequest(http.MethodPut, "/api/admin/fleets/"+id, strings.NewReader(reqBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Fleet not found", response["error"])

		mockService.AssertExpectations(t)
	})

	t.Run("Duplicate Plate Number", func(t *testing.T) {
		id := uuid.New().String()
		reqBody := `{"plate_number": "B 1234 XYZ"}`

		mockService.On("UpdateFleet", id, mock.AnythingOfType("*dto.UpdateFleetRequest")).Return(nil, apperrors.NewDuplicateField("plate_number")).Once()

		req, _ := http.NewRequest(http.MethodPut, "/api/admin/fleets/"+id, strings.NewReader(reqBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Validation failed", response["error"])

		mockService.AssertExpectations(t)
	})

	t.Run("Invalid Enum Type", func(t *testing.T) {
		id := uuid.New().String()
		reqBody := `{"type": "Hiace"}` // Invalid type, not Reguler or Premium

		req, _ := http.NewRequest(http.MethodPut, "/api/admin/fleets/"+id, strings.NewReader(reqBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Validation failed", response["error"])
	})

	t.Run("Service Error", func(t *testing.T) {
		id := uuid.New().String()
		reqBody := `{"name": "Armada 1 Updated"}`

		mockService.On("UpdateFleet", id, mock.AnythingOfType("*dto.UpdateFleetRequest")).Return(nil, errors.New("db error")).Once()

		req, _ := http.NewRequest(http.MethodPut, "/api/admin/fleets/"+id, strings.NewReader(reqBody))
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

func TestFleetController_DeleteFleet(t *testing.T) {
	mockService := new(MockFleetService)
	r := setupFleetRouter(mockService)

	t.Run("Success", func(t *testing.T) {
		id := uuid.New().String()
		mockService.On("DeleteFleet", id).Return(nil).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/api/admin/fleets/"+id, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Fleet deleted successfully", response["message"])

		mockService.AssertExpectations(t)
	})

	t.Run("Fleet Not Found", func(t *testing.T) {
		id := uuid.New().String()
		mockService.On("DeleteFleet", id).Return(apperrors.NewNotFound("Fleet")).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/api/admin/fleets/"+id, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Fleet not found", response["error"])

		mockService.AssertExpectations(t)
	})

	t.Run("Service Error", func(t *testing.T) {
		id := uuid.New().String()
		mockService.On("DeleteFleet", id).Return(errors.New("db error")).Once()

		req, _ := http.NewRequest(http.MethodDelete, "/api/admin/fleets/"+id, nil)
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
