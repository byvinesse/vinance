package service

import (
	"context"
	"time"

	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/cursor"
	"github.com/byvinesse/vinance-backend/repository"
)

const defaultRecordPageSize = 25

type RecordService struct {
	recordRepo      repository.Record
	recordLabelRepo repository.RecordLabel
}

func NewRecordService(recordRepo repository.Record, recordLabelRepo repository.RecordLabel) *RecordService {
	return &RecordService{
		recordRepo:      recordRepo,
		recordLabelRepo: recordLabelRepo,
	}
}

func (s *RecordService) CreateRecord(ctx context.Context, userID string, request *model.CreateRecordRequest) (*model.RecordResponse, error) {
	recordedAt := time.Now()
	if request.RecordedAt != nil {
		recordedAt = *request.RecordedAt
	}

	now := time.Now()
	payload := entity.Record{
		UserID:        userID,
		AccountID:     request.AccountID,
		SubCategoryID: request.SubCategoryID,
		Amount:        request.Amount,
		Currency:      request.Currency,
		BaseAmount:    request.BaseAmount,
		Type:          request.Type,
		Name:          request.Name,
		Payee:         request.Payee,
		PaymentType:   request.PaymentType,
		PaymentStatus: request.PaymentStatus,
		IsExcluded:    request.IsExcluded,
		RecordedAt:    recordedAt,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	res, err := s.recordRepo.InsertOne(ctx, &payload)
	if err != nil {
		return nil, err
	}

	if err := s.recordLabelRepo.InsertBatch(ctx, res.ID, request.Labels); err != nil {
		return nil, err
	}

	return toRecordResponse(res, request.Labels), nil
}

// GetRecords returns a paginated page of records for the given user.
// cursorStr is the opaque token from a previous response's next_cursor (empty = first page).
// limit <= 0 falls back to defaultRecordPageSize.
func (s *RecordService) GetRecords(ctx context.Context, userID string, limit int, cursorStr string) (*model.PaginatedRecordsResponse, error) {
	if limit <= 0 {
		limit = defaultRecordPageSize
	}

	decoded, err := cursor.Decode(cursorStr)
	if err != nil {
		return nil, err
	}

	var entityCursor *entity.RecordCursor
	if decoded != nil {
		entityCursor = &entity.RecordCursor{
			RecordedAt: decoded.RecordedAt,
			ID:         decoded.ID,
		}
	}

	records, err := s.recordRepo.FindByUserID(ctx, userID, limit, entityCursor)
	if err != nil {
		return nil, err
	}

	// Fetch all label associations for this page in a single query.
	labelsByRecordID, err := s.fetchLabelMap(ctx, records)
	if err != nil {
		return nil, err
	}

	responses := make([]model.RecordResponse, len(records))
	for i, r := range records {
		responses[i] = *toRecordResponse(&r, labelsByRecordID[r.ID])
	}

	// Build next_cursor from the last record in the page (nil when page is not full).
	var nextCursor *string
	if len(records) == limit {
		last := records[len(records)-1]
		encoded, err := cursor.Encode(cursor.RecordCursor{
			RecordedAt: last.RecordedAt,
			ID:         last.ID,
		})
		if err != nil {
			return nil, err
		}
		nextCursor = &encoded
	}

	return &model.PaginatedRecordsResponse{
		Records:    responses,
		NextCursor: nextCursor,
		Limit:      limit,
	}, nil
}

// fetchLabelMap returns a map of record ID → []labelID for all records in the page.
func (s *RecordService) fetchLabelMap(ctx context.Context, records []entity.Record) (map[string][]string, error) {
	if len(records) == 0 {
		return nil, nil
	}

	ids := make([]string, len(records))
	for i, r := range records {
		ids[i] = r.ID
	}

	rows, err := s.recordLabelRepo.FindByRecordIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	m := make(map[string][]string, len(records))
	for _, row := range rows {
		m[row.RecordID] = append(m[row.RecordID], row.LabelID)
	}
	return m, nil
}

func toRecordResponse(r *entity.Record, labels []string) *model.RecordResponse {
	if labels == nil {
		labels = []string{}
	}
	return &model.RecordResponse{
		ID:            r.ID,
		AccountID:     r.AccountID,
		SubCategoryID: r.SubCategoryID,
		Amount:        r.Amount,
		Currency:      r.Currency,
		BaseAmount:    r.BaseAmount,
		Type:          r.Type,
		Labels:        labels,
		Name:          r.Name,
		Payee:         r.Payee,
		PaymentType:   r.PaymentType,
		PaymentStatus: r.PaymentStatus,
		IsExcluded:    r.IsExcluded,
		RecordedAt:    r.RecordedAt,
		CreatedAt:     r.CreatedAt,
	}
}
