package repository

import (
	"context"

	"github.com/byvinesse/vinance-backend/entity"
)

type User interface {
	InsertOne(ctx context.Context, user *entity.User) (*entity.User, error)
	FindOneByEmail(ctx context.Context, email string) (*entity.User, error)
}

type Account interface {
	InsertOne(ctx context.Context, user *entity.Account) (*entity.Account, error)
}

type Category interface {
	FindCompleteCategory(ctx context.Context, userID string) ([]entity.CategoryWithSubCategory, error)
}
