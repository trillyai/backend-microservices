package service

import (
	"context"

	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/profile/contracts"
	"github.com/trillyai/backend-microservices/services/profile/shared"
)

type service struct {
	repository contracts.Repository
	logger     logger.Logger
}

func NewService(repository contracts.Repository) contracts.Service {
	return service{
		repository: repository,
		logger:     *logger.NewLogger("profile-service"),
	}
}

// GetProfileByUsername implements contracts.Service.
func (s service) GetProfileByUsername(ctx context.Context, username string) (shared.GetProfileResponse, error) {
	return s.repository.GetProfileByUsername(ctx, username)
}

// GetProfiles implements contracts.Service.
func (s service) GetProfiles(ctx context.Context, offset uint32, limit uint32) ([]shared.GetProfileResponse, error) {
	return s.repository.GetProfiles(ctx, offset, limit)
}

// UpdateProfile implements contracts.Service.
func (s service) UpdateProfile(ctx context.Context, request shared.UpdateProfileRequest) (shared.UpdateProfileResponse, error) {
	return s.repository.UpdateProfile(ctx, request)
}
