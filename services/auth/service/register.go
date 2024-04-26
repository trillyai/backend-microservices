package service

import (
	"context"

	"github.com/trillyai/backend-microservices/services/auth/shared"
)

func (s service) Register(ctx context.Context, request shared.RegisterRequest) (shared.RegisterResponse, error) {
	return s.repository.Register(ctx, request)
}
