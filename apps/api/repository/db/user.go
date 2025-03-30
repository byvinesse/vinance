package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/jmoiron/sqlx"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (r *User) InsertOne(ctx context.Context, user *entity.User) (*entity.User, error) {
	columns := []string{"email", "password_hash", "username", "phone_number", "gender", "created_at", "updated_at"}
	values := []interface{}{user.Email, user.PasswordHash, user.Username, user.PhoneNumber, user.Gender, user.CreatedAt, user.UpdatedAt}

	if user.DateOfBirth != nil {
		columns = append(columns, "date_of_birth")
		values = append(values, user.DateOfBirth)
	}

	queryBuilder := sq.Insert(entity.TableNameUser).
		Columns(columns...).
		Values(values...).
		Suffix("RETURNING *")

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var res entity.User
	err := r.db.GetContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *User) FindOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	queryBuilder := sq.Select("*").
		From(entity.TableNameUser).
		Where(sq.Eq{"email": email})

	query, args, _ := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()

	var user entity.User
	err := r.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
