package service

import (
	"context"

	"github.com/trillyai/backend-microservices/services/auth/shared"
)

func (s service) GetProfile(ctx context.Context, request shared.GetProfileRequest) (shared.GetProfileResponse, error) {
	return s.repository.GetProfile(ctx, request)
}
