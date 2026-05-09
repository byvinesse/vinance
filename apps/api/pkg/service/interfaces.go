package service

import (
	"context"

	"github.com/byvinesse/vinance-backend/model"
)

type IUserService interface {
	Register(ctx context.Context, request *model.RegisterRequest) (bool, error)
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
	GetProfile(ctx context.Context, email string) (*model.GetProfileResponse, error)
}

type IAccountService interface {
	CreateAccount(ctx context.Context, userID string, request *model.CreateAccountRequest) (bool, error)
}

type ICategoryService interface {
	GetCompleteCategory(ctx context.Context, userID string) ([]model.GetCompleteCategoriesResponse, error)
}

type IRecordService interface {
	CreateRecord(ctx context.Context, userID string, request *model.CreateRecordRequest) (*model.RecordResponse, error)
	GetRecords(ctx context.Context, userID string, limit int, cursor string) (*model.PaginatedRecordsResponse, error)
}
