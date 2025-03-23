package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/byvinesse/vinance-backend/config"
	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/authenticator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// MockUserRepository mocks the user repository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) InsertOne(ctx context.Context, user *entity.User) (*entity.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) FindOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestUserService_Register(t *testing.T) {
	// Setup
	mockRepo := new(MockUserRepository)
	testAuth := authenticator.NewAuthenticator(config.Auth{
		AccessTokenDuration: time.Hour,
		JwtKey:              "test-key",
	})

	service := NewUserService(mockRepo, *testAuth)

	// Test data
	ctx := context.Background()
	request := &model.RegisterRequest{
		Email:       "test@example.com",
		Password:    "password123",
		Username:    "testuser",
		PhoneNumber: "1234567890",
		Gender:      "male",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	t.Run("Success", func(t *testing.T) {
		// Setup expectations
		mockRepo.On("InsertOne", ctx, mock.AnythingOfType("*entity.User")).Return(&entity.User{
			ID:          "1",
			Email:       request.Email,
			Username:    request.Username,
			PhoneNumber: request.PhoneNumber,
			Gender:      request.Gender,
			DateOfBirth: request.DateOfBirth,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}, nil).Once()

		// Call the service
		result, err := service.Register(ctx, request)

		// Assert
		assert.NoError(t, err)
		assert.True(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Database Error", func(t *testing.T) {
		// Setup expectations
		mockRepo.On("InsertOne", ctx, mock.AnythingOfType("*entity.User")).Return(nil, errors.New("database error")).Once()

		// Call the service
		result, err := service.Register(ctx, request)

		// Assert
		assert.Error(t, err)
		assert.False(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_Login(t *testing.T) {
	// Setup
	mockRepo := new(MockUserRepository)
	testAuth := authenticator.NewAuthenticator(config.Auth{
		AccessTokenDuration: time.Hour,
		JwtKey:              "test-key",
	})

	service := NewUserService(mockRepo, *testAuth)

	// Test data
	ctx := context.Background()
	request := &model.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Generate a real hash for password123
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	user := &entity.User{
		ID:           "1",
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
		Username:     "testuser",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		// Setup expectations
		mockRepo.On("FindOneByEmail", ctx, request.Email).Return(user, nil).Once()

		// Call the service
		result, err := service.Login(ctx, request)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.AccessToken)
		mockRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		// Setup expectations
		mockRepo.On("FindOneByEmail", ctx, request.Email).Return(nil, errors.New("user not found")).Once()

		// Call the service
		result, err := service.Login(ctx, request)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Incorrect Password", func(t *testing.T) {
		// Create user with different password hash
		userWithDifferentPassword := *user
		incorrectHash, _ := bcrypt.GenerateFromPassword([]byte("different-password"), 14)
		userWithDifferentPassword.PasswordHash = string(incorrectHash)

		// Setup expectations
		mockRepo.On("FindOneByEmail", ctx, request.Email).Return(&userWithDifferentPassword, nil).Once()

		// Call the service
		result, err := service.Login(ctx, request)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}
