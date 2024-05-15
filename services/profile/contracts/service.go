package contracts

import (
	"context"

	"github.com/trillyai/backend-microservices/services/profile/shared"
)

type Service interface {
	GetProfiles(ctx context.Context, offset, limit uint32) ([]shared.GetProfileResponse, error)
	GetProfileByUsername(ctx context.Context, username string) (shared.GetProfileResponse, error)
	UpdateProfile(ctx context.Context, request shared.UpdateProfileRequest) (shared.UpdateProfileResponse, error)
	UploadProfileImage(ctx context.Context, request shared.UploadProfileImageRequest) (shared.UploadProfileImageResponse, error)
}
