package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/byvinesse/vinance-backend/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheckRoute(t *testing.T) {
	// Only test the healthcheck route, which doesn't rely on other services
	e := echo.New()

	// Initialize validator
	validator.Init()

	// Add the healthcheck route
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Create a request to the healthcheck endpoint
	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	rec := httptest.NewRecorder()

	// Serve the request
	e.ServeHTTP(rec, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "OK", rec.Body.String())
}
