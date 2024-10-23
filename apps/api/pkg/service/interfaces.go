package service

import (
	"context"

	"github.com/byvinesse/vinance-backend/model"
)

type IAuthService interface {
	Register(ctx context.Context, request *model.RegisterRequest) (bool, error)
	Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error)
	CompleteMemberOnboarding(ctx context.Context, request *model.CompleteMemberOnboardingRequest) (bool, error)
}

type IMemberService interface {
	CreateMember(ctx context.Context, request *model.CreateMemberRequest) (bool, error)
}
