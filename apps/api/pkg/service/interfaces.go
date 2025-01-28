package service

import (
	"context"

	"github.com/byvinesse/vinance-backend/model"
)

type IUserService interface {
	Register(ctx context.Context, request *model.RegisterRequest) (bool, error)
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
}
