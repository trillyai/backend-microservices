package repository

import (
	"context"

	"github.com/trillyai/backend-microservices/services/auth/shared"
)

func (r repository) Logout(ctx context.Context, request shared.LogoutRequest) (shared.LogoutResponse, error) {
	return shared.LogoutResponse{}, nil
}
