package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/byvinesse/vinance-backend/cmd/application"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRecordService struct {
	mock.Mock
}

func (m *MockRecordService) CreateRecord(ctx context.Context, userID string, request *model.CreateRecordRequest) (*model.RecordResponse, error) {
	args := m.Called(ctx, userID, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.RecordResponse), args.Error(1)
}

func (m *MockRecordService) GetRecords(ctx context.Context, userID string, limit int, cursor string) (*model.PaginatedRecordsResponse, error) {
	args := m.Called(ctx, userID, limit, cursor)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.PaginatedRecordsResponse), args.Error(1)
}

func TestHandler_CreateRecord(t *testing.T) {
	e := echo.New()
	mockRecordService := new(MockRecordService)

	app := &application.App{RecordService: mockRecordService}
	h := NewHandler(app)

	now := time.Now()
	validRequest := model.CreateRecordRequest{
		AccountID:     "acc-123",
		SubCategoryID: "subcat-456",
		Amount:        100,
		Currency:      entity.CurrencyTypeTHB,
		BaseAmount:    50000,
		Type:          entity.RecordTypeExpense,
		Labels:        []string{"label-1"},
		Name:          "Lunch",
		Payee:         "Restaurant",
		PaymentType:   entity.PaymentTypeCash,
		PaymentStatus: entity.PaymentStatusCleared,
		RecordedAt:    &now,
	}

	stubbedResponse := &model.RecordResponse{
		ID:            "rec-789",
		AccountID:     validRequest.AccountID,
		SubCategoryID: validRequest.SubCategoryID,
		Amount:        validRequest.Amount,
		Currency:      validRequest.Currency,
		BaseAmount:    validRequest.BaseAmount,
		Type:          validRequest.Type,
		Labels:        validRequest.Labels,
		Name:          validRequest.Name,
		Payee:         validRequest.Payee,
		PaymentType:   validRequest.PaymentType,
		PaymentStatus: validRequest.PaymentStatus,
		RecordedAt:    now,
		CreatedAt:     now,
	}

	newCtx := func(body interface{}) (echo.Context, *httptest.ResponseRecorder) {
		b, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/records/v1/_create", bytes.NewReader(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user_id", "test-user-id")
		c.Set("user_email", "test@example.com")
		return c, rec
	}

	t.Run("Success", func(t *testing.T) {
		mockRecordService.On("CreateRecord", mock.Anything, "test-user-id", mock.MatchedBy(func(req *model.CreateRecordRequest) bool {
			return req.AccountID == validRequest.AccountID &&
				req.Amount == validRequest.Amount &&
				req.Currency == validRequest.Currency &&
				req.BaseAmount == validRequest.BaseAmount
		})).Return(stubbedResponse, nil).Once()

		c, rec := newCtx(validRequest)
		if assert.NoError(t, h.CreateRecord(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			var resp entity.OkResponse[model.RecordResponse]
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
			assert.Equal(t, "rec-789", resp.Data.ID)
			assert.Equal(t, float64(100), resp.Data.Amount)
			assert.Equal(t, entity.CurrencyTypeTHB, resp.Data.Currency)
			assert.Equal(t, float64(50000), resp.Data.BaseAmount)
		}
		mockRecordService.AssertExpectations(t)
	})

	t.Run("Missing account_id", func(t *testing.T) {
		invalid := validRequest
		invalid.AccountID = ""
		c, _ := newCtx(invalid)
		assert.Error(t, h.CreateRecord(c))
	})

	t.Run("Missing subcategory_id", func(t *testing.T) {
		invalid := validRequest
		invalid.SubCategoryID = ""
		c, _ := newCtx(invalid)
		assert.Error(t, h.CreateRecord(c))
	})

	t.Run("Invalid amount", func(t *testing.T) {
		invalid := validRequest
		invalid.Amount = 0
		c, _ := newCtx(invalid)
		assert.Error(t, h.CreateRecord(c))
	})

	t.Run("Missing currency", func(t *testing.T) {
		invalid := validRequest
		invalid.Currency = ""
		c, _ := newCtx(invalid)
		assert.Error(t, h.CreateRecord(c))
	})

	t.Run("Invalid base_amount", func(t *testing.T) {
		invalid := validRequest
		invalid.BaseAmount = 0
		c, _ := newCtx(invalid)
		assert.Error(t, h.CreateRecord(c))
	})

	t.Run("Missing type", func(t *testing.T) {
		invalid := validRequest
		invalid.Type = ""
		c, _ := newCtx(invalid)
		assert.Error(t, h.CreateRecord(c))
	})

	t.Run("Missing payment_type", func(t *testing.T) {
		invalid := validRequest
		invalid.PaymentType = ""
		c, _ := newCtx(invalid)
		assert.Error(t, h.CreateRecord(c))
	})

	t.Run("Missing payment_status", func(t *testing.T) {
		invalid := validRequest
		invalid.PaymentStatus = ""
		c, _ := newCtx(invalid)
		assert.Error(t, h.CreateRecord(c))
	})
}

