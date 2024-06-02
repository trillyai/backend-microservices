package repository

import (
	"context"

	"github.com/trillyai/backend-microservices/services/profile/shared"
)

// GetUserInterests implements contracts.Repository.
func (r repository) GetUserInterests(ctx context.Context) ([]shared.Interest, error) {
	panic("unimplemented")
}

// GetInterests implements contracts.Repository.
func (r repository) GetInterests(ctx context.Context) ([]shared.Interest, error) {
	panic("unimplemented")
}

// CreateUserInterest implements contracts.Repository.
func (r repository) CreateUserInterest(ctx context.Context, request shared.CreateUserInterestRequest) (shared.CreateUserInterestResponse, error) {
	panic("unimplemented")
}

// DeleteUserInterest implements contracts.Repository.
func (r repository) DeleteUserInterest(ctx context.Context, request shared.DeleteUserInterestRequest) (shared.DeleteUserInterestResponse, error) {
	panic("unimplemented")
}
