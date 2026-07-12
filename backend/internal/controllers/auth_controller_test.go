package controllers_test

import (
	"bytes"
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

// Mock for AuthService
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(email, password string) (string, *models.Admin, error) {
	args := m.Called(email, password)
	if args.Get(1) != nil {
		return args.String(0), args.Get(1).(*models.Admin), args.Error(2)
	}
	return args.String(0), nil, args.Error(2)
}

func setupRouter(authService *MockAuthService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	authController := controllers.NewAuthController(authService)
	r.POST("/api/admin/login", authController.Login)
	return r
}

func TestAuthController_Login(t *testing.T) {
	mockService := new(MockAuthService)
	r := setupRouter(mockService)

	adminID := uuid.New()
	validAdmin := &models.Admin{
		ID:    adminID,
		Email: "admin@test.com",
		Name:  "Admin Test",
	}

	t.Run("Success Login", func(t *testing.T) {
		mockService.On("Login", "admin@test.com", "password123").Return("mocked-jwt-token", validAdmin, nil).Once()

		reqBody := map[string]string{
			"email":    "admin@test.com",
			"password": "password123",
		}
		jsonValue, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/api/admin/login", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Login successful", response["message"])
		
		data := response["data"].(map[string]interface{})
		assert.Equal(t, "mocked-jwt-token", data["token"])
		adminData := data["admin"].(map[string]interface{})
		assert.Equal(t, "admin@test.com", adminData["email"])

		mockService.AssertExpectations(t)
	})

	t.Run("Invalid Credentials", func(t *testing.T) {
		mockService.On("Login", "wrong@test.com", "wrongpass").Return("", nil, errors.New("invalid email or password")).Once()

		reqBody := map[string]string{
			"email":    "wrong@test.com",
			"password": "wrongpass",
		}
		jsonValue, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/api/admin/login", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "invalid email or password", response["error"])

		mockService.AssertExpectations(t)
	})

	t.Run("Invalid Payload (Missing Email)", func(t *testing.T) {
		reqBody := map[string]string{
			"password": "password123",
		}
		jsonValue, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, "/api/admin/login", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
