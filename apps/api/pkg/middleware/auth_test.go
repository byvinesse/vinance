package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/byvinesse/vinance-backend/config"
	"github.com/byvinesse/vinance-backend/entity"
	auth "github.com/byvinesse/vinance-backend/pkg/authenticator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	// Setup
	mockConfig := config.Auth{
		JwtKey:              "test-secret-key",
		AccessTokenDuration: time.Hour * 24,
	}
	authenticator := auth.NewAuthenticator(mockConfig)

	user := entity.User{
		ID:    "test-user-id",
		Email: "test@example.com",
	}

	token, _, err := authenticator.GenerateJwtToken(user)
	assert.NoError(t, err)

	// Test cases
	testCases := []struct {
		name           string
		authHeader     string
		expectedStatus int
		expectedUserID string
	}{
		{
			name:           "Valid token",
			authHeader:     "Bearer " + token,
			expectedStatus: http.StatusOK,
			expectedUserID: user.ID,
		},
		{
			name:           "Missing token",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			expectedUserID: "",
		},
		{
			name:           "Invalid token format",
			authHeader:     "Bearer invalid-token",
			expectedStatus: http.StatusUnauthorized,
			expectedUserID: "",
		},
		{
			name:           "Invalid header format",
			authHeader:     "InvalidFormat token",
			expectedStatus: http.StatusUnauthorized,
			expectedUserID: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup Echo
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			if tc.authHeader != "" {
				req.Header.Set(echo.HeaderAuthorization, tc.authHeader)
			}

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Create middleware
			authConfig := AuthConfig{
				AllowAuthorizationHeader: true,
			}
			middleware := Authentication(authenticator, authConfig)

			// Handler function that will be executed if middleware passes
			handlerFunc := func(c echo.Context) error {
				userID := c.Get("user_id")
				if userID != nil {
					return c.String(http.StatusOK, userID.(string))
				}
				return c.NoContent(http.StatusOK)
			}

			// Execute middleware
			err := middleware(handlerFunc)(c)

			// Assert
			if tc.expectedStatus == http.StatusOK {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedStatus, rec.Code)
				assert.Equal(t, tc.expectedUserID, rec.Body.String())
			} else {
				// If we expect an error, the middleware should return it
				assert.Error(t, err)
			}
		})
	}
}

func TestAuthenticateAuthorizationHeader(t *testing.T) {
	// Setup
	mockConfig := config.Auth{
		JwtKey:              "test-secret-key",
		AccessTokenDuration: time.Hour * 24,
	}
	authenticator := auth.NewAuthenticator(mockConfig)

	user := entity.User{
		ID:    "test-user-id",
		Email: "test@example.com",
	}

	token, _, err := authenticator.GenerateJwtToken(user)
	assert.NoError(t, err)

	// Test cases
	testCases := []struct {
		name           string
		authHeader     string
		expectError    bool
		expectedClaims *entity.TokenClaims
	}{
		{
			name:        "Valid token",
			authHeader:  "Bearer " + token,
			expectError: false,
			expectedClaims: &entity.TokenClaims{
				UserID:    user.ID,
				UserEmail: user.Email,
			},
		},
		{
			name:           "Missing header",
			authHeader:     "",
			expectError:    false,
			expectedClaims: nil,
		},
		{
			name:           "Invalid token format",
			authHeader:     "Bearer invalid-token",
			expectError:    true,
			expectedClaims: nil,
		},
		{
			name:           "Invalid header format",
			authHeader:     "InvalidFormat token",
			expectError:    true,
			expectedClaims: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup Echo
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			if tc.authHeader != "" {
				req.Header.Set(echo.HeaderAuthorization, tc.authHeader)
			}

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Execute function
			claims, err := authenticateAuthorizationHeader(c, authenticator)

			// Assert
			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, claims)
			} else {
				assert.NoError(t, err)
				if tc.expectedClaims == nil {
					assert.Nil(t, claims)
				} else {
					assert.NotNil(t, claims)
					assert.Equal(t, tc.expectedClaims.UserID, claims.UserID)
					assert.Equal(t, tc.expectedClaims.UserEmail, claims.UserEmail)
				}
			}
		})
	}
}
