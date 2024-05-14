package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// //////////////////////////////////////////////////////////////////////////////////
// CreateLike implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) CreateLike(ctx context.Context, req shared.CreateLikeRequest) (shared.CreateLikeResponse, error) {
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// DeleteLike implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) DeleteLike(ctx context.Context, req shared.DeleteLikeRequest) (shared.DeleteLikeResponse, error) {
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// GetLikes implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) GetLikes(ctx context.Context, uuid uuid.UUID, forPostId bool, forUserId bool, offset uint32, limit uint32) (shared.CreateLikeResponse, error) {
	panic("unimplemented")
}
