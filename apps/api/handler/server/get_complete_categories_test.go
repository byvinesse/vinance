package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/byvinesse/vinance-backend/cmd/application"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) GetCompleteCategory(ctx context.Context, userID string) ([]model.GetCompleteCategoriesResponse, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.GetCompleteCategoriesResponse), args.Error(1)
}

func TestHandler_GetCompleteCategories(t *testing.T) {
	// Setup
	e := echo.New()
	mockCategoryService := new(MockCategoryService)

	app := &application.App{
		CategoryService: mockCategoryService,
	}

	h := NewHandler(app)

	// Initialize Validator
	validator.Init()

	// Create test request data
	userID := "test-user-id"

	t.Run("Success", func(t *testing.T) {
		// Create expected response
		expected := []model.GetCompleteCategoriesResponse{
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
		}

		// Mock service response
		mockCategoryService.On("GetCompleteCategory", mock.Anything, userID).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/categories/v1/complete", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Set user_id in context
		c.Set("user_id", userID)

		// Test
		if assert.NoError(t, h.GetCompleteCategories(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			// Parse response
			var response entity.OkResponse[[]model.GetCompleteCategoriesResponse]
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, 200, response.Code)
			assert.Equal(t, "OK", response.Status)

			// Assert on category data
			assert.Equal(t, expected[0], response.Data[0])
			assert.Equal(t, expected[1], response.Data[1])
		}

		// Verify mocks
		mockCategoryService.AssertExpectations(t)
	})
}
