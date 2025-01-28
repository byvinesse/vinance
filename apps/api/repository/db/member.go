package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/vincentkdeli/vinance-backend/entity"
)

type Member struct {
	db *sqlx.DB
}

func NewMember(db *sqlx.DB) *Member {
	return &Member{
		db: db,
	}
}

func (r *Member) InsertOne(ctx context.Context, member *entity.Member) (*entity.Member, error) {
	queryBuilder := sq.Insert(entity.TableNameMember).
		Columns("email", "username", "phone_number", "date_of_birth", "gender", "created_at", "updated_at").
		Values(member.Email, member.Username, member.PhoneNumber, member.DateOfBirth, member.Gender, member.CreatedAt, member.UpdatedAt).
		Suffix("RETURNING *")

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var res entity.Member
	err := r.db.GetContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
