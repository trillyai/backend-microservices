package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/services/auth/shared"
	"golang.org/x/crypto/bcrypt"
)

func (r repository) Register(ctx context.Context, request shared.RegisterRequest) (shared.RegisterResponse, error) {
	r.logger.Debug(fmt.Sprintf("register request recived with Username: %s", request.Username))

	request.Username = strings.ToLower(request.Username)
	readResp, err := postgres.Read[tables.User, tables.User](ctx, map[string]interface{}{"Username": request.Username})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.RegisterResponse{}, err
	}

	if readResp.Username == request.Username {
		return shared.RegisterResponse{}, errors.New("the Username has already been taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.RegisterResponse{}, err
	}

	request.Password = string(hashedPassword)
	user, err := postgres.Create[shared.RegisterResponse, tables.User](ctx, request)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.RegisterResponse{}, err
	}

	return user, nil
}
