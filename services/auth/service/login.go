package service

import (
	"context"

	"github.com/trillyai/backend-microservices/services/auth/shared"
)

func (s service) Login(ctx context.Context, request shared.LoginRequest) (shared.LoginResponse, error) {

	return s.repository.Login(ctx, request)
}
