package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) InsertOne(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	args := m.Called(ctx, account)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestAccountService_CreateAccount(t *testing.T) {
	// Setup
	mockRepo := new(MockAccountRepository)

	service := NewAccountService(mockRepo)

	// Test data
	ctx := context.Background()
	request := &model.CreateAccountRequest{
		Name:     "Test Account",
		Balance:  1000,
		Currency: entity.CurrencyTypeIDR,
		Type:     entity.AccountTypeCash,
		Color:    "#000",
	}

	t.Run("Success", func(t *testing.T) {
		// Setup expectations
		mockRepo.On("InsertOne", ctx, mock.AnythingOfType("*entity.Account")).Return(&entity.Account{
			ID:            "1",
			UserID:        "test-user-id",
			Name:          "Test Account",
			Balance:       1000,
			Currency:      entity.CurrencyTypeIDR,
			Type:          entity.AccountTypeCash,
			Color:         "#000",
			IsArchived:    false,
			IsExcluded:    false,
			MarkForDelete: false,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}, nil).Once()

		// Call the service
		result, err := service.CreateAccount(ctx, "test-user-id", request)

		// Assert
		assert.NoError(t, err)
		assert.True(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		// Setup expectations
		mockRepo.On("InsertOne", ctx, mock.AnythingOfType("*entity.Account")).Return(nil, errors.New("database error")).Once()

		// Call the service
		result, err := service.CreateAccount(ctx, "test-uesr-id", request)

		// Assert
		assert.Error(t, err)
		assert.False(t, result)
		mockRepo.AssertExpectations(t)
	})

}
