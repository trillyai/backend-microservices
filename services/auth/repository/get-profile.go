package repository

import (
	"context"
	"errors"

	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/services/auth/shared"
)

func (r repository) GetProfile(ctx context.Context, request shared.GetProfileRequest) (shared.GetProfileResponse, error) {

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.GetProfileResponse{}, errors.New("context not found")
	}

	resp, err := postgres.Read[shared.GetProfileResponse, tables.User](ctx, map[string]interface{}{"Id": claims.Id})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.GetProfileResponse{}, err
	}

	return resp, nil
}
