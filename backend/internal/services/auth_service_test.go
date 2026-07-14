package services_test

import (
	"errors"
	"testing"

	"mesenshuttle-backend/internal/config"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/services"
	"mesenshuttle-backend/pkg/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Mock for AdminRepository
type MockAdminRepository struct {
	mock.Mock
}

func (m *MockAdminRepository) FindByEmail(email string) (*models.Admin, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Admin), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestAuthService_Login(t *testing.T) {
	mockRepo := new(MockAdminRepository)
	cfg := &config.Config{
		JWTSecret: "testsecret",
	}

	authService := services.NewAuthService(mockRepo, cfg)

	hashPwd := func(password string, cost int) string {
		hash, _ := utils.HashPassword(password, cost)
		return hash
	}
	adminID := uuid.New()
	validAdmin := &models.Admin{
		ID:           adminID,
		Email:        "admin@test.com",
		PasswordHash: hashPwd("password123", 14),
		Name:         "Admin Test",
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("FindByEmail", "admin@test.com").Return(validAdmin, nil).Once()

		token, admin, err := authService.Login("admin@test.com", "password123")

		assert.NoError(t, err)
		assert.NotNil(t, admin)
		assert.NotEmpty(t, token)
		assert.Equal(t, "admin@test.com", admin.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Email (Not Found)", func(t *testing.T) {
		mockRepo.On("FindByEmail", "wrong@test.com").Return(nil, gorm.ErrRecordNotFound).Once()

		token, admin, err := authService.Login("wrong@test.com", "password123")

		assert.Error(t, err)
		assert.Equal(t, "invalid email or password", err.Error())
		assert.Nil(t, admin)
		assert.Empty(t, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		mockRepo.On("FindByEmail", "admin@test.com").Return(validAdmin, nil).Once()

		token, admin, err := authService.Login("admin@test.com", "wrongpassword")

		assert.Error(t, err)
		assert.Equal(t, "invalid email or password", err.Error())
		assert.Nil(t, admin)
		assert.Empty(t, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		mockRepo.On("FindByEmail", "error@test.com").Return(nil, errors.New("db error")).Once()

		token, admin, err := authService.Login("error@test.com", "password123")

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		assert.Nil(t, admin)
		assert.Empty(t, token)
		mockRepo.AssertExpectations(t)
	})
}
