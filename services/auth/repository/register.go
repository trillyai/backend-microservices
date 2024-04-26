package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/services/auth/shared"
)

func (r repository) Register(ctx context.Context, request shared.RegisterRequest) (shared.RegisterResponse, error) {
	r.logger.Debug(fmt.Sprintf("register request recived with username: %s", request.UserName))

	readResp, err := postgres.Read[tables.User, tables.User](ctx, map[string]interface{}{"UserName": request.UserName})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.RegisterResponse{}, err
	}

	if readResp.UserName == request.UserName {
		return shared.RegisterResponse{}, errors.New("the username has already been taken")
	}

	user, err := postgres.Create[shared.RegisterResponse, tables.User](ctx, request)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.RegisterResponse{}, err
	}

	return user, nil
}
