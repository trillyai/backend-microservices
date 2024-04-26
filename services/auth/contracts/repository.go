package contracts

import (
	"context"

	"github.com/trillyai/backend-microservices/services/auth/shared"
)

type Repository interface {
	Register(ctx context.Context, request shared.RegisterRequest) (shared.RegisterResponse, error)
}
