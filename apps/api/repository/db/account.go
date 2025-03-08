package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/jmoiron/sqlx"
)

type Account struct {
	db *sqlx.DB
}

func NewAccount(db *sqlx.DB) *Account {
	return &Account{
		db: db,
	}
}

func (r *Account) InsertOne(ctx context.Context, user *entity.Account) (*entity.Account, error) {
	queryBuilder := sq.Insert(entity.TableNameAccount).
		Columns("user_id", "name", "balance", "currency", "type", "color", "is_archived", "is_excluded", "mark_for_delete", "created_at", "updated_at").
		Values(user.UserID, user.Name, user.Balance, user.Currency, user.Type, user.Color, user.IsArchived, user.IsExcluded, user.MarkForDelete, user.CreatedAt, user.UpdatedAt).
		Suffix("RETURNING *")

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var res entity.Account
	err := r.db.GetContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
