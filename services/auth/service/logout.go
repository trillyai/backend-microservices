package service

import (
	"context"

	"github.com/trillyai/backend-microservices/services/auth/shared"
)

func (s service) Logout(ctx context.Context, request shared.LogoutRequest) (shared.LogoutResponse, error) {
	return s.repository.Logout(ctx, request)
}
