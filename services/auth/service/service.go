package service

import (
	"context"

	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/auth/contracts"
	"github.com/trillyai/backend-microservices/services/auth/shared"
)

type service struct {
	repository contracts.Repository
	logger     logger.Logger
}

func NewService(repository contracts.Repository) contracts.Service {
	return service{
		repository: repository,
		logger:     *logger.NewLogger("auth-service"),
	}
}

// //////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////
func (s service) Register(ctx context.Context, request shared.RegisterRequest) (shared.RegisterResponse, error) {
	return s.repository.Register(ctx, request)
}

// //////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////
func (s service) Logout(ctx context.Context, request shared.LogoutRequest) (shared.LogoutResponse, error) {
	return s.repository.Logout(ctx, request)
}

// //////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////
func (s service) Login(ctx context.Context, request shared.LoginRequest) (shared.LoginResponse, error) {
	return s.repository.Login(ctx, request)
}
