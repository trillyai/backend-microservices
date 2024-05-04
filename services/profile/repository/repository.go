package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
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
	resp, err := postgres.Read[shared.GetProfileResponse, tables.User](ctx, map[string]interface{}{"Username": username})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.GetProfileResponse{}, err
	}

	if resp.Id == uuid.Nil {
		return shared.GetProfileResponse{}, errors.New("user not found")
	}

	return resp, nil
}

// GetProfiles implements contracts.Repository.
func (r repository) GetProfiles(ctx context.Context, offset uint32, limit uint32) ([]shared.GetProfileResponse, error) {
	panic("unimplemented")
}

// UpdateProfile implements contracts.Repository.
func (r repository) UpdateProfile(ctx context.Context, request shared.UpdateProfileRequest) (shared.UpdateProfileResponse, error) {
	claims := ctx.Value("user").(*auth.Claims)
	if claims.UserName == "" {
		return shared.UpdateProfileResponse{}, errors.New("context not found")
	}
	panic("unimplemented")
}
