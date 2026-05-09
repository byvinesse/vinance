package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/jmoiron/sqlx"
)

type Record struct {
	db *sqlx.DB
}

func NewRecord(db *sqlx.DB) *Record {
	return &Record{db: db}
}

func (r *Record) InsertOne(ctx context.Context, record *entity.Record) (*entity.Record, error) {
	queryBuilder := sq.Insert(entity.TableNameRecord).
		Columns(
			"user_id", "account_id", "subcategory_id",
			"amount", "currency", "base_amount",
			"type", "name", "payee",
			"payment_type", "payment_status",
			"is_excluded", "recorded_at", "created_at", "updated_at",
		).
		Values(
			record.UserID, record.AccountID, record.SubCategoryID,
			record.Amount, record.Currency, record.BaseAmount,
			record.Type, record.Name, record.Payee,
			record.PaymentType, record.PaymentStatus,
			record.IsExcluded, record.RecordedAt, record.CreatedAt, record.UpdatedAt,
		).
		Suffix("RETURNING *")

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var res entity.Record
	if err := r.db.GetContext(ctx, &res, query, args...); err != nil {
		return nil, err
	}

	return &res, nil
}

// FindByUserID retrieves records for a user with cursor-based keyset pagination.
// Records are ordered by recorded_at DESC, id DESC.
// When cursor is non-nil, only records strictly before the cursor position are returned.
func (r *Record) FindByUserID(ctx context.Context, userID string, limit int, cursor *entity.RecordCursor) ([]entity.Record, error) {
	queryBuilder := sq.Select("*").
		From(entity.TableNameRecord).
		OrderBy("recorded_at DESC, id DESC").
		Limit(uint64(limit))

	if cursor == nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"user_id": userID})
	} else {
		// Keyset: next page starts strictly after (recorded_at, id) of the last seen row.
		// (recorded_at < cursor) OR (recorded_at = cursor AND id < cursor_id)
		queryBuilder = queryBuilder.Where(
			sq.And{
				sq.Eq{"user_id": userID},
				sq.Or{
					sq.Lt{"recorded_at": cursor.RecordedAt},
					sq.And{
						sq.Eq{"recorded_at": cursor.RecordedAt},
						sq.Lt{"id": cursor.ID},
					},
				},
			},
		)
	}

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var res []entity.Record
	if err := r.db.SelectContext(ctx, &res, query, args...); err != nil {
		return nil, err
	}

	return res, nil
}
