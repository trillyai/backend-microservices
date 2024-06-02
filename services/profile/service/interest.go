package service

import (
	"context"

	"github.com/trillyai/backend-microservices/services/profile/shared"
)

// GetUserInterests implements contracts.Service.
func (s service) GetUserInterests(ctx context.Context) ([]shared.Interest, error) {
	return s.repository.GetInterests(ctx)
}

// GetInterests implements contracts.Service.
func (s service) GetInterests(ctx context.Context) ([]shared.Interest, error) {
	return s.repository.GetInterests(ctx)
}

// CreateUserInterest implements contracts.Service.
func (s service) CreateUserInterest(ctx context.Context, request shared.CreateUserInterestRequest) (shared.CreateUserInterestResponse, error) {
	return s.repository.CreateUserInterest(ctx, request)
}

// DeleteUserInterest implements contracts.Service.
func (s service) DeleteUserInterest(ctx context.Context, request shared.DeleteUserInterestRequest) (shared.DeleteUserInterestResponse, error) {
	return s.repository.DeleteUserInterest(ctx, request)
}
