package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// //////////////////////////////////////////////////////////////////////////////////
// CreateLike implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) CreateLike(ctx context.Context, req shared.CreateLikeRequest) (shared.CreateLikeResponse, error) {
	return s.repository.CreateLike(ctx, req)
}

// //////////////////////////////////////////////////////////////////////////////////
// DeleteLike implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) DeleteLike(ctx context.Context, req shared.DeleteLikeRequest) (shared.DeleteLikeResponse, error) {
	return s.repository.DeleteLike(ctx, req)
}

// //////////////////////////////////////////////////////////////////////////////////
// GetLikes implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) GetLikes(ctx context.Context, uuid uuid.UUID, forPostId bool, forCommentId bool, offset uint32, limit uint32) (shared.Likes, error) {
	return s.repository.GetLikes(ctx, uuid, forPostId, forCommentId, offset, limit)
}
