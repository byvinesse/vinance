package repository

import (
	"context"

	"github.com/vincentkdeli/vinance-backend/entity"
)

type Auth interface {
	InsertOne(ctx context.Context, auth *entity.Auth) (*entity.Auth, error)
	FindOneByEmail(ctx context.Context, email string) (*entity.Auth, error)
	UpdateOne(ctx context.Context, id string) (*entity.Auth, error)
}

type Member interface {
	InsertOne(ctx context.Context, member *entity.Member) (*entity.Member, error)
}
