package authenticator

import (
	"testing"
	"time"

	"github.com/byvinesse/vinance-backend/config"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticator_GenerateJwtToken(t *testing.T) {
	// Setup
	mockConfig := config.Auth{
		JwtKey:              "test-secret-key",
		AccessTokenDuration: time.Hour * 24,
	}
	authenticator := NewAuthenticator(mockConfig)

	user := entity.User{
		ID:    "test-user-id",
		Email: "test@example.com",
	}

	// Execute
	token, expiresAt, err := authenticator.GenerateJwtToken(user)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.WithinDuration(t, time.Now().Add(mockConfig.AccessTokenDuration), expiresAt, time.Second*5)

	// Verify token can be parsed
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(mockConfig.JwtKey), nil
	})

	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, user.ID, claims["user_id"])
	assert.Equal(t, user.Email, claims["user_email"])
}

func TestAuthenticator_ParseToken(t *testing.T) {
	// Setup
	mockConfig := config.Auth{
		JwtKey:              "test-secret-key",
		AccessTokenDuration: time.Hour * 24,
	}
	authenticator := NewAuthenticator(mockConfig)

	user := entity.User{
		ID:    "test-user-id",
		Email: "test@example.com",
	}

	token, _, err := authenticator.GenerateJwtToken(user)
	assert.NoError(t, err)

	// Test case 1: Valid token
	t.Run("Valid token", func(t *testing.T) {
		// Execute
		claims, err := authenticator.ParseToken(token)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, claims)
		assert.Equal(t, user.ID, claims.UserID)
		assert.Equal(t, user.Email, claims.UserEmail)
	})

	// Test case 2: Invalid token
	t.Run("Invalid token", func(t *testing.T) {
		// Execute
		claims, err := authenticator.ParseToken("invalid.token.string")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, claims)
	})

	// Test case 3: Expired token
	t.Run("Expired token", func(t *testing.T) {
		// Create a token that's already expired
		expiredConfig := config.Auth{
			JwtKey:              "test-secret-key",
			AccessTokenDuration: -time.Hour, // Negative duration to make it expired
		}
		expiredAuthenticator := NewAuthenticator(expiredConfig)

		expiredToken, _, err := expiredAuthenticator.GenerateJwtToken(user)
		assert.NoError(t, err)

		// Execute
		claims, err := authenticator.ParseToken(expiredToken)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, claims)
	})
}
