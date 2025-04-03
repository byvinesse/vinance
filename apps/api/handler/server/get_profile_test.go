package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/byvinesse/vinance-backend/cmd/application"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (m *MockUserService) GetProfile(ctx context.Context, email string) (*model.GetProfileResponse, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.GetProfileResponse), args.Error(1)
}

func TestHandler_GetProfile(t *testing.T) {
	// Setup
	e := echo.New()
	mockUserService := new(MockUserService)

	app := &application.App{
		UserService: mockUserService,
	}

	h := NewHandler(app)

	// Initialize Validator
	validator.Init()

	// Create test request data
	email := "test@example.com"

	t.Run("Success", func(t *testing.T) {
		// Create expected profile response
		expected := &model.GetProfileResponse{
			Email:       email,
			Username:    "testuser",
			PhoneNumber: "1234567890",
			Gender:      "M",
			DateOfBirth: time.Now().AddDate(-30, 0, 0).UnixMilli(), // 30 years ago
		}

		// Mock service response
		mockUserService.On("GetProfile", mock.Anything, email).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/users/v1/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Set user_email in context
		c.Set("user_email", email)

		// Test
		if assert.NoError(t, h.GetProfile(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			// Parse response
			var response entity.OkResponse[model.GetProfileResponse]
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, 200, response.Code)
			assert.Equal(t, "OK", response.Status)

			// Assert on profile data
			assert.Equal(t, expected.Email, response.Data.Email)
			assert.Equal(t, expected.Username, response.Data.Username)
			assert.Equal(t, expected.PhoneNumber, response.Data.PhoneNumber)
			assert.Equal(t, expected.Gender, response.Data.Gender)
			assert.Equal(t, expected.DateOfBirth, response.Data.DateOfBirth)
		}
	})

	// Verify mocks
	mockUserService.AssertExpectations(t)
}
