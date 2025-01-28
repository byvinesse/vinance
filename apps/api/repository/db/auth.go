package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/vincentkdeli/vinance-backend/entity"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (r *Auth) InsertOne(ctx context.Context, auth *entity.Auth) (*entity.Auth, error) {
	queryBuilder := sq.Insert(entity.TableNameAuth).
		Columns("email", "password", "is_member", "created_at", "updated_at").
		Values(auth.Email, auth.Password, auth.IsMember, auth.CreatedAt, auth.UpdatedAt).
		Suffix("RETURNING *")

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var res entity.Auth
	err := r.db.GetContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Auth) FindOneByEmail(ctx context.Context, email string) (*entity.Auth, error) {
	queryBuilder := sq.Select("*").
		From(entity.TableNameAuth).
		Where(sq.Eq{"email": email})

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var auth entity.Auth
	err := r.db.GetContext(ctx, &auth, query, args...)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *Auth) UpdateOne(ctx context.Context, id string) (*entity.Auth, error) {
	queryBuilder := sq.Update(entity.TableNameAuth).
		Set("is_member", true).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING *")

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var res entity.Auth
	err := r.db.QueryRowxContext(ctx, query, args...).StructScan(&res)
	if err != nil {
		return nil, err
	}

	return &res, err
}
