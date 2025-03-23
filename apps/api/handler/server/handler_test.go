package server

import (
	"testing"

	"github.com/byvinesse/vinance-backend/cmd/application"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	// Create a mock App
	mockApp := &application.App{}

	// Initialize handler with mock app
	handler := NewHandler(mockApp)

	// Assert handler is created correctly
	assert.NotNil(t, handler)
	assert.Equal(t, mockApp, handler.app)
}
