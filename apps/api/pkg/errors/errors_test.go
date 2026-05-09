package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorTypes(t *testing.T) {
	t.Run("validation errors", func(t *testing.T) {
		err := ErrMissingField("username")
		assert.Equal(t, 400, err.Code)
		assert.Equal(t, "MISSING_FIELD", err.Status)
		assert.Equal(t, "username field cannot be empty", err.Message)
		assert.Equal(t, "username field cannot be empty", err.Error())

		err2 := ErrInvalidValue("password")
		assert.Equal(t, 400, err2.Code)
		assert.Equal(t, "INVALID_VALUE", err2.Status)
		assert.Equal(t, "password field value is invalid", err2.Message)
		assert.Equal(t, "password field value is invalid", err2.Error())

		err3 := ErrInvalidFormat("email", "email address")
		assert.Equal(t, 400, err3.Code)
		assert.Equal(t, "INVALID_FORMAT", err3.Status)
		assert.Equal(t, "email field value is not a valid email address", err3.Message)
		assert.Equal(t, "email field value is not a valid email address", err3.Error())

		err4 := ErrMissingPathParam("id")
		assert.Equal(t, 400, err4.Code)
		assert.Equal(t, "MISSING_PARAMETER", err4.Status)
		assert.Equal(t, "id path parameter cannot be empty", err4.Message)
		assert.Equal(t, "id path parameter cannot be empty", err4.Error())

		internalErr := errors.New("json parsing failed")
		err5 := ErrParseFailed(internalErr)
		assert.Equal(t, 400, err5.Code)
		assert.Equal(t, "PARSE_ERROR", err5.Status)
		assert.Equal(t, "Failed to parse request body", err5.Message)
		assert.Equal(t, "Failed to parse request body", err5.Error())
		assert.Equal(t, internalErr, err5.InternalError)
	})

	t.Run("unauthorized error", func(t *testing.T) {
		internalErr := errors.New("invalid token")
		err := ErrUnauthorized(internalErr, "Authentication failed")
		assert.Equal(t, 401, err.Code)
		assert.Equal(t, "UNAUTHORIZED", err.Status)
		assert.Equal(t, "Authentication failed", err.Message)
		assert.Equal(t, "Authentication failed", err.Error())
		assert.Equal(t, internalErr, err.InternalError)
	})

	t.Run("forbidden error", func(t *testing.T) {
		internalErr := errors.New("permission denied")
		err := ErrForbidden(internalErr, "Insufficient permissions")
		assert.Equal(t, 403, err.Code)
		assert.Equal(t, "FORBIDDEN", err.Status)
		assert.Equal(t, "Insufficient permissions", err.Message)
		assert.Equal(t, "Insufficient permissions", err.Error())
		assert.Equal(t, internalErr, err.InternalError)
	})

	t.Run("not found error", func(t *testing.T) {
		internalErr := errors.New("record not found")
		err := ErrDataNotFoundError(internalErr, "User not found")
		assert.Equal(t, 404, err.Code)
		assert.Equal(t, "DATA_NOT_FOUND", err.Status)
		assert.Equal(t, "User not found", err.Message)
		assert.Equal(t, "User not found", err.Error())
		assert.Equal(t, internalErr, err.InternalError)
	})

	t.Run("duplicate error", func(t *testing.T) {
		internalErr := errors.New("duplicate entry")
		err := ErrDuplicateError(internalErr, "Email already exists")
		assert.Equal(t, 409, err.Code)
		assert.Equal(t, "DUPLICATE_ERROR", err.Status)
		assert.Equal(t, "Email already exists", err.Message)
		assert.Equal(t, "Email already exists", err.Error())
		assert.Equal(t, internalErr, err.InternalError)
	})

	t.Run("server errors", func(t *testing.T) {
		internalErr := errors.New("server crashed")
		err := ErrInternalServerError(internalErr, "Internal server error")
		assert.Equal(t, 500, err.Code)
		assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Status)
		assert.Equal(t, "Internal server error", err.Message)
		assert.Equal(t, "Internal server error", err.Error())
		assert.Equal(t, internalErr, err.InternalError)

		dbErr := errors.New("connection timeout")
		err2 := DatabaseError(dbErr, "Database connection failed")
		assert.Equal(t, 500, err2.Code)
		assert.Equal(t, "DATABASE_ERROR", err2.Status)
		assert.Equal(t, "Database connection failed", err2.Message)
		assert.Equal(t, "Database connection failed", err2.Error())
		assert.Equal(t, dbErr, err2.InternalError)
	})
}
