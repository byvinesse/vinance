package service

import "github.com/vincentkdeli/vinance-backend/repository"

type AuthService struct {
	authRepo repository.Auth
}

func NewAuthService(authRepo repository.Auth) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}
