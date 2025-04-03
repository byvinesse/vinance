package application

import (
	"context"
	"testing"

	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/pkg/authenticator"
	"github.com/byvinesse/vinance-backend/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock implementations
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(ctx context.Context, request *model.RegisterRequest) (bool, error) {
	args := m.Called(ctx, request)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserService) Login(ctx context.Context, request *model.LoginRequest) (*model.LoginResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*model.LoginResponse), args.Error(1)
}

func (m *MockUserService) GetProfile(ctx context.Context, email string) (*model.GetProfileResponse, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.GetProfileResponse), args.Error(1)
}

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) CreateAccount(ctx context.Context, userID string, request *model.CreateAccountRequest) (bool, error) {
	args := m.Called(ctx, userID, request)
	return args.Bool(0), args.Error(1)
}

func TestApp_Initialization(t *testing.T) {
	// Create mock services
	mockUserService := new(MockUserService)
	mockAccountService := new(MockAccountService)
	mockAuthenticator := new(authenticator.Authenticator)

	// Create app with mocks
	app := &App{
		UserService:    mockUserService,
		AccountService: mockAccountService,
		Authenticator:  mockAuthenticator,
	}

	// Verify that the app has been initialized with the correct components
	assert.NotNil(t, app)
	assert.Equal(t, mockUserService, app.UserService)
	assert.Equal(t, mockAccountService, app.AccountService)
	assert.Equal(t, mockAuthenticator, app.Authenticator)
}

// This test verifies that the app struct satisfies the expected interfaces
func TestApp_InterfaceCompliance(t *testing.T) {
	// Verify that our mock implementations satisfy the interfaces
	var userService service.IUserService = new(MockUserService)
	var accountService service.IAccountService = new(MockAccountService)

	// Simple assertion to ensure the interfaces are satisfied
	assert.NotNil(t, userService)
	assert.NotNil(t, accountService)
}
