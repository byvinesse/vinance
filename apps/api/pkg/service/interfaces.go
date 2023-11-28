package service

import (
	"context"

	"github.com/vincentkdeli/vinance-backend/model"
)

type IAuthService interface {
	Register(ctx context.Context, request *model.RegisterRequest) (*model.RegisterResponse, error)
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
}
