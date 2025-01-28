package service

import (
	"context"
	"fmt"
	"time"

	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	auth "github.com/byvinesse/vinance-backend/pkg/authenticator"
	"github.com/byvinesse/vinance-backend/pkg/errors"
	"github.com/byvinesse/vinance-backend/pkg/utils"
	"github.com/byvinesse/vinance-backend/repository"
)

type UserService struct {
	userRepo      repository.User
	authenticator auth.Authenticator
}

func NewUserService(userRepo repository.User, authenticator auth.Authenticator) *UserService {
	return &UserService{
		userRepo:      userRepo,
		authenticator: authenticator,
	}
}

func (s *UserService) Register(ctx context.Context, request *model.RegisterRequest) (bool, error) {
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return false, err
	}

	payload := entity.User{
		Email:        request.Email,
		PasswordHash: hashedPassword,
		Username:     request.Username,
		PhoneNumber:  request.PhoneNumber,
		DateOfBirth:  request.DateOfBirth,
		Gender:       request.Gender,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	res, err := s.userRepo.InsertOne(ctx, &payload)
	if err != nil {
		return false, err
	}

	return res != nil, nil
}

func (s *UserService) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	data, err := s.userRepo.FindOneByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(request.Password, data.PasswordHash) {
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
