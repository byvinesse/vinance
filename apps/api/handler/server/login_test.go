package server

import (
	"bytes"
	"encoding/json"
	stderrors "errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/byvinesse/vinance-backend/cmd/application"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Login(t *testing.T) {
	// Setup
	e := echo.New()
	mockUserService := new(MockUserService)

	app := &application.App{
		UserService: mockUserService,
	}

	h := NewHandler(app)

	// Test data
	validRequest := model.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	validResponse := &model.LoginResponse{
		AccessToken:          "dummy-token",
		AccessTokenExpiresAt: time.Now().Add(time.Hour),
	}

	t.Run("Success", func(t *testing.T) {
		// Setup mock
		mockUserService.On("Login", mock.Anything, mock.MatchedBy(func(req *model.LoginRequest) bool {
			return req.Email == validRequest.Email && req.Password == validRequest.Password
		})).Return(validResponse, nil).Once()

		// Create request
		requestBody, _ := json.Marshal(validRequest)
		req := httptest.NewRequest(http.MethodPost, "/auth/v1/login", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Test
		if assert.NoError(t, h.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			// Parse response
			var response entity.OkResponse[model.LoginResponse]
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, 200, response.Code)
			assert.Equal(t, "OK", response.Status)
			assert.Equal(t, validResponse.AccessToken, response.Data.AccessToken)
		}

		// Verify all expectations were met
		mockUserService.AssertExpectations(t)
	})

	t.Run("Missing Email", func(t *testing.T) {
		// Create request with missing email
		invalidRequest := validRequest
		invalidRequest.Email = ""

		requestBody, _ := json.Marshal(invalidRequest)
		req := httptest.NewRequest(http.MethodPost, "/auth/v1/login", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Test
		err := h.Login(c)
		assert.Error(t, err)
	})

	t.Run("Missing Password", func(t *testing.T) {
		// Create request with missing password
		invalidRequest := validRequest
		invalidRequest.Password = ""

		requestBody, _ := json.Marshal(invalidRequest)
		req := httptest.NewRequest(http.MethodPost, "/auth/v1/login", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Test
		err := h.Login(c)
		assert.Error(t, err)
	})

	t.Run("Invalid Credentials", func(t *testing.T) {
		// Setup mock to return unauthorized error
		mockUserService.On("Login", mock.Anything, mock.MatchedBy(func(req *model.LoginRequest) bool {
			return req.Email == validRequest.Email && req.Password == validRequest.Password
		})).Return(nil, errors.ErrUnauthorized(stderrors.New("INCORRECT_PASSWORD"), "sorry, incorrect password")).Once()

		// Create request
		requestBody, _ := json.Marshal(validRequest)
		req := httptest.NewRequest(http.MethodPost, "/auth/v1/login", bytes.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Test
		err := h.Login(c)
		assert.Error(t, err)

		// Verify all expectations were met
		mockUserService.AssertExpectations(t)
	})
}
