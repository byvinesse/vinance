package repository

import (
	"context"

	"github.com/vincentkdeli/vinance-backend/entity"
)

type Auth interface {
	InsertOne(ctx context.Context, auth *entity.Auth) (*entity.Auth, error)
	FindOneByEmail(ctx context.Context, email string) (*entity.Auth, error)
}
