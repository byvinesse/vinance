package service

import (
	"context"
	"fmt"
	"time"

	"github.com/vincentkdeli/vinance-backend/entity"
	"github.com/vincentkdeli/vinance-backend/model"
	auth "github.com/vincentkdeli/vinance-backend/pkg/authenticator"
	"github.com/vincentkdeli/vinance-backend/pkg/errors"
	"github.com/vincentkdeli/vinance-backend/pkg/utils"
	"github.com/vincentkdeli/vinance-backend/repository"
)

type AuthService struct {
	authRepo      repository.Auth
	authenticator auth.Authenticator
}

func NewAuthService(authRepo repository.Auth, authenticator auth.Authenticator) *AuthService {
	return &AuthService{
		authRepo:      authRepo,
		authenticator: authenticator,
	}
}

func (s *AuthService) Register(ctx context.Context, request *model.RegisterRequest) (*model.RegisterResponse, error) {
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	payload := entity.Auth{
		Email:     request.Email,
		Password:  hashedPassword,
		IsMember:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := s.authRepo.InsertOne(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return toRegisterResponse(res), nil
}

func toRegisterResponse(auth *entity.Auth) *model.RegisterResponse {
	return &model.RegisterResponse{
		Email:    auth.Email,
		IsMember: auth.IsMember,
	}
}

func (s *AuthService) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	data, err := s.authRepo.FindOneByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(request.Password, data.Password) {
		return nil, errors.ErrUnauthorized(fmt.Errorf("INCORRECT_PASSWORD"), "sorry, incorrect password")
	}

	accessToken, accessTokenExpiresAt, err := s.authenticator.GenerateJwtToken(*data)
	if err != nil {
		return nil, err
	}

	return toLoginResponse(accessToken, accessTokenExpiresAt), nil
}

func toLoginResponse(accessToken string, accessTokenExpiresAt time.Time) *model.LoginResponse {
	return &model.LoginResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessTokenExpiresAt,
	}
}
