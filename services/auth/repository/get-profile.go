package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/services/auth/shared"
)

func (r repository) GetProfile(ctx context.Context, request shared.GetProfileRequest) (shared.GetProfileResponse, error) {
	claims := ctx.Value("user").(*auth.Claims)
	if claims.UserName == "" {
		return shared.GetProfileResponse{}, errors.New("context not found")
	}

	resp, err := postgres.Read[shared.GetProfileResponse, tables.User](ctx, map[string]interface{}{"Id": claims.UserId})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.GetProfileResponse{}, err
	}

	if resp.Id == uuid.Nil {
		return shared.GetProfileResponse{}, errors.New("user not found")
	}

	return resp, nil
}
