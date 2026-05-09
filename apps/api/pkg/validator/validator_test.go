package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=130"`
}

func TestValidateStruct(t *testing.T) {
	// Initialize validator
	Init()

	ctx := context.Background()

	t.Run("valid struct", func(t *testing.T) {
		validStruct := TestStruct{
			Name:  "John Doe",
			Email: "john@example.com",
			Age:   30,
		}

		err := ValidateStruct(ctx, validStruct)
		assert.NoError(t, err)
	})

	t.Run("missing required field", func(t *testing.T) {
		invalidStruct := TestStruct{
			Email: "john@example.com",
			Age:   30,
		}

		err := ValidateStruct(ctx, invalidStruct)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name")
	})

	t.Run("invalid email format", func(t *testing.T) {
		invalidStruct := TestStruct{
			Name:  "John Doe",
			Email: "not-an-email",
			Age:   30,
		}

		err := ValidateStruct(ctx, invalidStruct)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email")
	})

	t.Run("age out of range", func(t *testing.T) {
		invalidStruct := TestStruct{
			Name:  "John Doe",
			Email: "john@example.com",
			Age:   150,
		}

		err := ValidateStruct(ctx, invalidStruct)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "age")
	})
}
