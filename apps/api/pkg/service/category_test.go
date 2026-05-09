package service

import (
	"context"
	"errors"
	"testing"

	"github.com/byvinesse/vinance-backend/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) FindCompleteCategory(ctx context.Context, userID string) ([]entity.CategoryWithSubCategory, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.CategoryWithSubCategory), args.Error(1)
}

func TestCategoryService_GetCompleteCategory(t *testing.T) {
	// Setup
	mockRepo := new(MockCategoryRepository)

	service := NewCategoryService(mockRepo)

	// Test data
	ctx := context.Background()
	userID := "test-user-id"

	t.Run("Success", func(t *testing.T) {
		// Setup expectations
		mockRepo.On("FindCompleteCategory", ctx, userID).Return([]entity.CategoryWithSubCategory{
			{
				CategoryID:      "1",
				SubCategoryID:   "1",
				CategoryName:    "Category 1",
				SubCategoryName: "SubCategory 1",
			},
			{
				CategoryID:      "2",
				SubCategoryID:   "2",
				CategoryName:    "Category 2",
				SubCategoryName: "SubCategory 2",
			},
		}, nil).Once()

		// Call the service
		result, err := service.GetCompleteCategory(ctx, userID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		// Setup expectations
		mockRepo.On("FindCompleteCategory", ctx, userID).Return(nil, errors.New("database error")).Once()

		// Call the service
		result, err := service.GetCompleteCategory(ctx, userID)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}
