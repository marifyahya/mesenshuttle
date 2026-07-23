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

type MockScheduleRepository struct {
	mock.Mock
}

func (m *MockScheduleRepository) FindAll(page, limit int) ([]models.Schedule, int64, error) {
	args := m.Called(page, limit)
	if args.Get(0) != nil {
		return args.Get(0).([]models.Schedule), args.Get(1).(int64), args.Error(2)
	}
	return nil, 0, args.Error(2)
}

func TestScheduleService_GetAllSchedules(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	scheduleService := services.NewScheduleService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		mockSchedules := []models.Schedule{
			{
				ID:      uuid.New(),
				RouteID: uuid.New(),
				FleetID: uuid.New(),
				Price:   150000,
				Status:  "Scheduled",
			},
		}
		mockRepo.On("FindAll", 1, 10).Return(mockSchedules, int64(1), nil).Once()

		schedules, total, err := scheduleService.GetAllSchedules(1, 10)

		assert.NoError(t, err)
		assert.NotNil(t, schedules)
		assert.Equal(t, int64(1), total)
		assert.Len(t, schedules, 1)
		assert.Equal(t, float64(150000), schedules[0].Price)
		assert.Equal(t, "Scheduled", schedules[0].Status)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		mockRepo.On("FindAll", 1, 10).Return(nil, int64(0), errors.New("db error")).Once()

		schedules, total, err := scheduleService.GetAllSchedules(1, 10)

		assert.Error(t, err)
		assert.Equal(t, "db error", err.Error())
		assert.Nil(t, schedules)
		assert.Equal(t, int64(0), total)
		mockRepo.AssertExpectations(t)
	})
}
