package repository

import (
	"context"

	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/profile/contracts"
	"github.com/trillyai/backend-microservices/services/profile/shared"
)

type repository struct {
	logger logger.Logger
}

func NewRepository() contracts.Repository {
	return repository{
		logger: *logger.NewLogger("profile-repository"),
	}
}

// GetProfileByUsername implements contracts.Repository.
func (r repository) GetProfileByUsername(ctx context.Context, username string) (shared.GetProfileResponse, error) {
	panic("unimplemented")
}

// GetProfiles implements contracts.Repository.
func (r repository) GetProfiles(ctx context.Context, offset uint32, limit uint32) ([]shared.GetProfileResponse, error) {
	panic("unimplemented")
}

// UpdateProfile implements contracts.Repository.
func (r repository) UpdateProfile(ctx context.Context, request shared.UpdateProfileRequest) (shared.UpdateProfileResponse, error) {
	panic("unimplemented")
}
