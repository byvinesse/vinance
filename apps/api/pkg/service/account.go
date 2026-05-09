package service

import (
	"context"
	"time"

	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/repository"
)

type AccountService struct {
	accountRepo repository.Account
}

func NewAccountService(accountRepo repository.Account) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

func (s *AccountService) CreateAccount(ctx context.Context, userID string, request *model.CreateAccountRequest) (bool, error) {
	payload := entity.Account{
		UserID:        userID,
		Name:          request.Name,
		Balance:       request.Balance,
		Currency:      request.Currency,
		Type:          request.Type,
		Color:         "#000",
		IsArchived:    false,
		IsExcluded:    false,
		MarkForDelete: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	res, err := s.accountRepo.InsertOne(ctx, &payload)
	if err != nil {
		return false, err
	}

	return res != nil, nil
}
