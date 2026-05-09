package server

import (
	"encoding/json"
	"errors"
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

// newGetRecordsCtx builds an echo.Context for GET /records/v1 with the given query string.
func newGetRecordsCtx(e *echo.Echo, query string) (echo.Context, *httptest.ResponseRecorder) {
	target := "/records/v1"
	if query != "" {
		target += "?" + query
	}
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user_id", "user-abc")
	c.Set("user_email", "test@example.com")
	return c, rec
}

// sampleRecord returns a fully-populated RecordResponse for use in stubs.
func sampleRecord(id string, now time.Time) model.RecordResponse {
	return model.RecordResponse{
		ID:            id,
		AccountID:     "acc-001",
		SubCategoryID: "subcat-001",
		Amount:        100,
		Currency:      entity.CurrencyTypeTHB,
		BaseAmount:    50000,
		Type:          entity.RecordTypeExpense,
		Labels:        []string{"label-1", "label-2"},
		Name:          "Lunch",
		Payee:         "Restaurant ABC",
		PaymentType:   entity.PaymentTypeCash,
		PaymentStatus: entity.PaymentStatusCleared,
		IsExcluded:    false,
		RecordedAt:    now,
		CreatedAt:     now,
	}
}

func TestHandler_GetRecords(t *testing.T) {
	e := echo.New()
	now := time.Now().Truncate(time.Second)

	// -------------------------------------------------------------------------
	// Happy-path cases
	// -------------------------------------------------------------------------

	t.Run("Success/FirstPage_NoParams", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		stub := &model.PaginatedRecordsResponse{
			Records: []model.RecordResponse{sampleRecord("rec-1", now)},
			Limit:   25,
		}
		svc.On("GetRecords", mock.Anything, "user-abc", 0, "").Return(stub, nil).Once()

		c, rec := newGetRecordsCtx(e, "")
		if assert.NoError(t, h.GetRecords(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var resp entity.OkResponse[model.PaginatedRecordsResponse]
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))

			assert.Equal(t, 200, resp.Code)
			assert.Equal(t, "OK", resp.Status)
			assert.Len(t, resp.Data.Records, 1)
			assert.Equal(t, 25, resp.Data.Limit)
			assert.Nil(t, resp.Data.NextCursor)
		}
		svc.AssertExpectations(t)
	})

	t.Run("Success/MultipleRecords_FullFieldAssertion", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		rec1 := sampleRecord("rec-1", now)
		rec2 := model.RecordResponse{
			ID:            "rec-2",
			AccountID:     "acc-002",
			SubCategoryID: "subcat-002",
			Amount:        200,
			Currency:      entity.CurrencyTypeIDR,
			BaseAmount:    200,
			Type:          entity.RecordTypeIncome,
			Labels:        []string{},
			Name:          "Salary",
			Payee:         "Company",
			PaymentType:   entity.PaymentTypeBankTransfer,
			PaymentStatus: entity.PaymentStatusCleared,
			IsExcluded:    false,
			RecordedAt:    now,
			CreatedAt:     now,
		}
		stub := &model.PaginatedRecordsResponse{
			Records: []model.RecordResponse{rec1, rec2},
			Limit:   25,
		}
		svc.On("GetRecords", mock.Anything, "user-abc", 0, "").Return(stub, nil).Once()

		c, rec := newGetRecordsCtx(e, "")
		if assert.NoError(t, h.GetRecords(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var resp entity.OkResponse[model.PaginatedRecordsResponse]
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
			assert.Len(t, resp.Data.Records, 2)

			// First record — verify all fields
			r1 := resp.Data.Records[0]
			assert.Equal(t, "rec-1", r1.ID)
			assert.Equal(t, "acc-001", r1.AccountID)
			assert.Equal(t, "subcat-001", r1.SubCategoryID)
			assert.Equal(t, float64(100), r1.Amount)
			assert.Equal(t, entity.CurrencyTypeTHB, r1.Currency)
			assert.Equal(t, float64(50000), r1.BaseAmount)
			assert.Equal(t, entity.RecordTypeExpense, r1.Type)
			assert.Equal(t, []string{"label-1", "label-2"}, r1.Labels)
			assert.Equal(t, "Lunch", r1.Name)
			assert.Equal(t, "Restaurant ABC", r1.Payee)
			assert.Equal(t, entity.PaymentTypeCash, r1.PaymentType)
			assert.Equal(t, entity.PaymentStatusCleared, r1.PaymentStatus)
			assert.False(t, r1.IsExcluded)

			// Second record — spot-check differing fields
			r2 := resp.Data.Records[1]
			assert.Equal(t, "rec-2", r2.ID)
			assert.Equal(t, entity.CurrencyTypeIDR, r2.Currency)
			assert.Equal(t, entity.RecordTypeIncome, r2.Type)
			assert.Empty(t, r2.Labels)
		}
		svc.AssertExpectations(t)
	})

	t.Run("Success/WithNextCursor_FullPage", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		cursor := "eyJyZWNvcmRlZF9hdCI6IjIwMjYtMDEtMDFUMDA6MDA6MDBaIiwiaWQiOiJyZWMtMjUifQ=="
		stub := &model.PaginatedRecordsResponse{
			Records:    []model.RecordResponse{sampleRecord("rec-25", now)},
			NextCursor: &cursor,
			Limit:      25,
		}
		svc.On("GetRecords", mock.Anything, "user-abc", 0, "").Return(stub, nil).Once()

		c, rec := newGetRecordsCtx(e, "")
		if assert.NoError(t, h.GetRecords(c)) {
			var resp entity.OkResponse[model.PaginatedRecordsResponse]
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
			assert.NotNil(t, resp.Data.NextCursor)
			assert.Equal(t, cursor, *resp.Data.NextCursor)
		}
		svc.AssertExpectations(t)
	})

	t.Run("Success/LastPage_NilNextCursor", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		stub := &model.PaginatedRecordsResponse{
			Records:    []model.RecordResponse{sampleRecord("rec-last", now)},
			NextCursor: nil,
			Limit:      25,
		}
		svc.On("GetRecords", mock.Anything, "user-abc", 0, "").Return(stub, nil).Once()

		c, rec := newGetRecordsCtx(e, "")
		if assert.NoError(t, h.GetRecords(c)) {
			var resp entity.OkResponse[model.PaginatedRecordsResponse]
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
			assert.Nil(t, resp.Data.NextCursor)
		}
		svc.AssertExpectations(t)
	})

	t.Run("Success/CursorAndLimit_BothProvided", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		stub := &model.PaginatedRecordsResponse{Records: []model.RecordResponse{}, Limit: 10}
		svc.On("GetRecords", mock.Anything, "user-abc", 10, "some-cursor").Return(stub, nil).Once()

		c, rec := newGetRecordsCtx(e, "cursor=some-cursor&limit=10")
		if assert.NoError(t, h.GetRecords(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var resp entity.OkResponse[model.PaginatedRecordsResponse]
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
			assert.Equal(t, 10, resp.Data.Limit)
		}
		svc.AssertExpectations(t)
	})

	t.Run("Success/CursorOnly_DefaultLimit", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		stub := &model.PaginatedRecordsResponse{Records: []model.RecordResponse{}, Limit: 25}
		svc.On("GetRecords", mock.Anything, "user-abc", 0, "my-cursor").Return(stub, nil).Once()

		c, _ := newGetRecordsCtx(e, "cursor=my-cursor")
		assert.NoError(t, h.GetRecords(c))
		svc.AssertExpectations(t)
	})

	t.Run("Success/LimitOnly_NoCursor", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		stub := &model.PaginatedRecordsResponse{Records: []model.RecordResponse{}, Limit: 5}
		svc.On("GetRecords", mock.Anything, "user-abc", 5, "").Return(stub, nil).Once()

		c, _ := newGetRecordsCtx(e, "limit=5")
		assert.NoError(t, h.GetRecords(c))
		svc.AssertExpectations(t)
	})

	t.Run("Success/EmptyRecords", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		stub := &model.PaginatedRecordsResponse{Records: []model.RecordResponse{}, Limit: 25}
		svc.On("GetRecords", mock.Anything, "user-abc", 0, "").Return(stub, nil).Once()

		c, rec := newGetRecordsCtx(e, "")
		if assert.NoError(t, h.GetRecords(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var resp entity.OkResponse[model.PaginatedRecordsResponse]
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
			assert.Empty(t, resp.Data.Records)
			assert.Nil(t, resp.Data.NextCursor)
		}
		svc.AssertExpectations(t)
	})

	// -------------------------------------------------------------------------
	// Input validation
	// -------------------------------------------------------------------------

	t.Run("InvalidLimit/NonNumeric", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		c, _ := newGetRecordsCtx(e, "limit=abc")
		err := h.GetRecords(c)
		assert.Error(t, err)
		svc.AssertNotCalled(t, "GetRecords")
	})

	t.Run("InvalidLimit/Zero", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		c, _ := newGetRecordsCtx(e, "limit=0")
		err := h.GetRecords(c)
		assert.Error(t, err)
		svc.AssertNotCalled(t, "GetRecords")
	})

	t.Run("InvalidLimit/Negative", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		c, _ := newGetRecordsCtx(e, "limit=-5")
		err := h.GetRecords(c)
		assert.Error(t, err)
		svc.AssertNotCalled(t, "GetRecords")
	})

	// -------------------------------------------------------------------------
	// Service error
	// -------------------------------------------------------------------------

	t.Run("ServiceError_Returns500", func(t *testing.T) {
		svc := new(MockRecordService)
		h := NewHandler(&application.App{RecordService: svc})

		svc.On("GetRecords", mock.Anything, "user-abc", 0, "").
			Return(nil, errors.New("db connection lost")).Once()

		c, _ := newGetRecordsCtx(e, "")
		err := h.GetRecords(c)
		assert.Error(t, err)
		svc.AssertExpectations(t)
	})
}
