package repository

import (
	"context"

	"github.com/byvinesse/vinance-backend/entity"
)

type User interface {
	InsertOne(ctx context.Context, user *entity.User) (*entity.User, error)
	FindOneByEmail(ctx context.Context, email string) (*entity.User, error)
}
