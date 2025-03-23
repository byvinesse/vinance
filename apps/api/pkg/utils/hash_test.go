package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	// Test case 1: Valid password
	t.Run("Valid password", func(t *testing.T) {
		// Setup
		password := "secure_password123"

		// Execute
		hashedPassword, err := HashPassword(password)

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, hashedPassword)
		assert.NotEqual(t, password, hashedPassword)

		// Verify that we can check the password against the hash
		assert.True(t, CheckPasswordHash(password, hashedPassword))
	})

	// Test case 2: Empty password
	t.Run("Empty password", func(t *testing.T) {
		// Setup
		password := ""

		// Execute
		hashedPassword, err := HashPassword(password)

		// Assert
		assert.NoError(t, err) // bcrypt can hash empty strings
		assert.NotEmpty(t, hashedPassword)

		// Verify that we can check the password against the hash
		assert.True(t, CheckPasswordHash(password, hashedPassword))
	})
}

func TestCheckPasswordHash(t *testing.T) {
	// Setup
	password := "secure_password123"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)

	// Test case 1: Correct password
	t.Run("Correct password", func(t *testing.T) {
		// Execute & Assert
		assert.True(t, CheckPasswordHash(password, hashedPassword))
	})

	// Test case 2: Incorrect password
	t.Run("Incorrect password", func(t *testing.T) {
		// Execute & Assert
		assert.False(t, CheckPasswordHash("wrong_password", hashedPassword))
	})

	// Test case 3: Empty password
	t.Run("Empty password", func(t *testing.T) {
		// Execute & Assert
		assert.False(t, CheckPasswordHash("", hashedPassword))
	})

	// Test case 4: Empty hash
	t.Run("Empty hash", func(t *testing.T) {
		// Execute & Assert
		assert.False(t, CheckPasswordHash(password, ""))
	})

	// Test case 5: Invalid hash format
	t.Run("Invalid hash format", func(t *testing.T) {
		// Execute & Assert
		assert.False(t, CheckPasswordHash(password, "not_a_valid_bcrypt_hash"))
	})
}
