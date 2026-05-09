package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/jmoiron/sqlx"
)

type RecordLabel struct {
	db *sqlx.DB
}

func NewRecordLabel(db *sqlx.DB) *RecordLabel {
	return &RecordLabel{db: db}
}

// InsertBatch inserts all label associations for a single record in one query.
// It is a no-op when labelIDs is empty.
func (r *RecordLabel) InsertBatch(ctx context.Context, recordID string, labelIDs []string) error {
	if len(labelIDs) == 0 {
		return nil
	}

	queryBuilder := sq.Insert(entity.TableNameRecordLabel).
		Columns("record_id", "label_id")

	for _, id := range labelIDs {
		queryBuilder = queryBuilder.Values(recordID, id)
	}

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

// FindByRecordIDs returns all label associations for the given record IDs.
func (r *RecordLabel) FindByRecordIDs(ctx context.Context, recordIDs []string) ([]entity.RecordLabel, error) {
	if len(recordIDs) == 0 {
		return nil, nil
	}

	queryBuilder := sq.Select("record_id", "label_id").
		From(entity.TableNameRecordLabel).
		Where(sq.Eq{"record_id": recordIDs})

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var res []entity.RecordLabel
	if err := r.db.SelectContext(ctx, &res, query, args...); err != nil {
		return nil, err
	}

	return res, nil
}
