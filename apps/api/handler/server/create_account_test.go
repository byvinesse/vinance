package server

import (
	"bytes"
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

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) CreateAccount(ctx context.Context, userID string, request *model.CreateAccountRequest) (bool, error) {
	args := m.Called(ctx, userID, request)
	return args.Bool(0), args.Error(1)
}

func TestHandler_CreateAccount(t *testing.T) {
	// Setup
	e := echo.New()
	mockAccountService := new(MockAccountService)

	app := &application.App{
		AccountService: mockAccountService,
	}

	h := NewHandler(app)

	// Initialize Validator
	validator.Init()

	// Create test request data
	validRequest := model.CreateAccountRequest{
		Name:     "Test Account",
		Balance:  1000.0,
		Currency: entity.CurrencyTypeIDR,
		Type:     entity.AccountTypeCash,
		Color:    "#000",
	}

	t.Run("Success", func(t *testing.T) {
		// Mock service response
		mockAccountService.On("CreateAccount", mock.Anything, mock.AnythingOfType("string"), mock.MatchedBy(func(req *model.CreateAccountRequest) bool {
			return req.Name == validRequest.Name
		})).Return(true, nil).Once()

		// Create request
		requestBody, _ := json.Marshal(validRequest)
		req := httptest.NewRequest(http.MethodPost, "/accounts/v1/_create", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Set user_id in context
		c.Set("user_id", "test-user-id")
		c.Set("user_email", "test@example.com")

		// Test
		if assert.NoError(t, h.CreateAccount(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			// Parse response
			var response entity.OkResponse[bool]
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, 201, response.Code)
			assert.Equal(t, "Created", response.Status)
			assert.True(t, response.Data)
		}

		// Verify mocks
		mockAccountService.AssertExpectations(t)
	})

	t.Run("Missing Name", func(t *testing.T) {
		// Create request with missing name
		invalidRequest := validRequest
		invalidRequest.Name = ""

		requestBody, _ := json.Marshal(invalidRequest)
		req := httptest.NewRequest(http.MethodPost, "/accounts/v1/_create", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Set user_id in context
		c.Set("user_id", "test-user-id")
		c.Set("user_email", "test@example.com")

		// Test - should return error
		err := h.CreateAccount(c)
		assert.Error(t, err)
	})

	t.Run("Missing Balance", func(t *testing.T) {
		// Create request with missing balance
		request := model.CreateAccountRequest{
			Name:     "Without Balance",
			Currency: entity.CurrencyTypeIDR,
			Type:     entity.AccountTypeCash,
			Color:    "#000",
		}

		// Mock service response
		mockAccountService.On("CreateAccount", mock.Anything, mock.AnythingOfType("string"), mock.MatchedBy(func(req *model.CreateAccountRequest) bool {
			return req.Name == request.Name
		})).Return(true, nil).Once()

		// Create request
		requestBody, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/accounts/v1/_create", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Set user_id in context
		c.Set("user_id", "test-user-id")
		c.Set("user_email", "test@example.com")

		// Test
		if assert.NoError(t, h.CreateAccount(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			// Parse response
			var response entity.OkResponse[bool]
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, 201, response.Code)
			assert.Equal(t, "Created", response.Status)
			assert.True(t, response.Data)
		}

		// Verify mocks
		mockAccountService.AssertExpectations(t)
	})

	t.Run("Missing Currency", func(t *testing.T) {
		// Create request with missing currency
		invalidRequest := validRequest
		invalidRequest.Currency = ""

		requestBody, _ := json.Marshal(invalidRequest)
		req := httptest.NewRequest(http.MethodPost, "/accounts/v1/_create", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Set user_id in context
		c.Set("user_id", "test-user-id")
		c.Set("user_email", "test@example.com")

		// Test - should return error
		err := h.CreateAccount(c)
		assert.Error(t, err)
	})

	t.Run("Missing Type", func(t *testing.T) {
		// Create request with missing type
		invalidRequest := validRequest
		invalidRequest.Type = ""

		requestBody, _ := json.Marshal(invalidRequest)
		req := httptest.NewRequest(http.MethodPost, "/accounts/v1/_create", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Set user_id in context
		c.Set("user_id", "test-user-id")
		c.Set("user_email", "test@example.com")

		// Test - should return error
		err := h.CreateAccount(c)
		assert.Error(t, err)
	})

	t.Run("Missing Color", func(t *testing.T) {
		// Create request with missing color
		invalidRequest := validRequest
		invalidRequest.Color = ""

		requestBody, _ := json.Marshal(invalidRequest)
		req := httptest.NewRequest(http.MethodPost, "/accounts/v1/_create", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Set user_id in context
		c.Set("user_id", "test-user-id")
		c.Set("user_email", "test@example.com")

		// Test - should return error
		err := h.CreateAccount(c)
		assert.Error(t, err)
	})

}
