package contracts

import (
	"context"

	"github.com/trillyai/backend-microservices/services/auth/shared"
)

type Service interface {
	Register(ctx context.Context, request shared.RegisterRequest) (shared.RegisterResponse, error)
	Login(ctx context.Context, request shared.LoginRequest) (shared.LoginResponse, error)
	Logout(ctx context.Context, request shared.LogoutRequest) (shared.LogoutResponse, error)
	GetProfile(ctx context.Context, request shared.GetProfileRequest) (shared.GetProfileResponse, error)
}
