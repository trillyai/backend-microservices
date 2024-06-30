package contracts

import (
	"context"

	"github.com/trillyai/backend-microservices/services/profile/shared"
)

type Repository interface {
	GetProfiles(ctx context.Context, offset, limit uint32) ([]shared.GetProfileResponse, error)
	GetProfileByUsername(ctx context.Context, username string) (shared.GetProfileResponse, error)
	UpdateProfile(ctx context.Context, request shared.UpdateProfileRequest) (shared.UpdateProfileResponse, error)
	UploadProfileImage(ctx context.Context, request shared.UploadProfileImageRequest) (shared.UploadProfileImageResponse, error)

	GetUserInterests(ctx context.Context, username string) ([]shared.Interest, error)
	CreateUserInterest(ctx context.Context, request shared.CreateUserInterestRequest) (shared.CreateUserInterestResponse, error)
	DeleteUserInterest(ctx context.Context, request shared.DeleteUserInterestRequest) (shared.DeleteUserInterestResponse, error)
	GetInterests(ctx context.Context) ([]shared.Interest, error) // metadata
}
