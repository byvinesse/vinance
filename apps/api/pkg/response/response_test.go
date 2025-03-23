package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/byvinesse/vinance-backend/entity"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestOk(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	testData := TestData{
		ID:   "123",
		Name: "Test User",
	}

	// Act
	err := Ok(c, testData)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBody entity.OkResponse[TestData]
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, 200, responseBody.Code)
	assert.Equal(t, "OK", responseBody.Status)
	assert.Equal(t, testData, responseBody.Data)
}

func TestOkCreated(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	testData := TestData{
		ID:   "456",
		Name: "New User",
	}

	// Act
	err := OkCreated(c, testData)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var responseBody entity.OkResponse[TestData]
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, 201, responseBody.Code)
	assert.Equal(t, "Created", responseBody.Status)
	assert.Equal(t, testData, responseBody.Data)
}
