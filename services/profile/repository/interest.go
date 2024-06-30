package repository

import (
	"context"

	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/services/profile/shared"
)

// GetUserInterests implements contracts.Repository.
func (r repository) GetUserInterests(ctx context.Context, username string) ([]shared.Interest, error) {
	resp, err := postgres.Read[[]shared.Interest, tables.Interest](ctx, map[string]interface{}{})
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
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
