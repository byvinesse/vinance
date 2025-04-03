package server

import (
	"bytes"
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

// Mock UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(ctx context.Context, request *model.RegisterRequest) (bool, error) {
	args := m.Called(ctx, request)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserService) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.LoginResponse), args.Error(1)
}

func TestHandler_Register(t *testing.T) {
	// Setup
	e := echo.New()
	mockUserService := new(MockUserService)

	app := &application.App{
		UserService: mockUserService,
	}

	h := NewHandler(app)

	// Initialize validator
	validator.Init()

	// Create test request data
	validRequest := model.RegisterRequest{
		Email:       "test@example.com",
		Password:    "password123",
		Username:    "testuser",
		PhoneNumber: "1234567890",
		Gender:      "M",
		DateOfBirth: func() *time.Time {
			t := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
			return &t
		}(),
	}

	t.Run("Success", func(t *testing.T) {
		// Mock service response
		mockUserService.On("Register", mock.Anything, mock.MatchedBy(func(req *model.RegisterRequest) bool {
			return req.Email == validRequest.Email
		})).Return(true, nil).Once()

		// Create request
		requestBody, _ := json.Marshal(validRequest)
		req := httptest.NewRequest(http.MethodPost, "/auth/v1/register", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Test
		if assert.NoError(t, h.Register(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			// Parse response
			var response entity.OkResponse[bool]
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, 200, response.Code)
			assert.Equal(t, "OK", response.Status)
			assert.True(t, response.Data)
		}

		// Verify mocks
		mockUserService.AssertExpectations(t)
	})

	t.Run("Missing Email", func(t *testing.T) {
		// Create request with missing email
		invalidRequest := validRequest
		invalidRequest.Email = ""

		requestBody, _ := json.Marshal(invalidRequest)
		req := httptest.NewRequest(http.MethodPost, "/auth/v1/register", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Test - should return error
		err := h.Register(c)
		assert.Error(t, err)
	})

	t.Run("Missing Password", func(t *testing.T) {
		// Create request with missing password
		invalidRequest := validRequest
		invalidRequest.Password = ""

		requestBody, _ := json.Marshal(invalidRequest)
		req := httptest.NewRequest(http.MethodPost, "/auth/v1/register", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Test - should return error
		err := h.Register(c)
		assert.Error(t, err)
	})
}
