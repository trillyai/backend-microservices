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
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// DeleteLike implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) DeleteLike(ctx context.Context, req shared.DeleteLikeRequest) (shared.DeleteLikeReesponse, error) {
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// GetLikes implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) GetLikes(ctx context.Context, uuid uuid.UUID, forPostId bool, forUserId bool, offset uint32, limit uint32) ([]shared.Like, error) {
	panic("unimplemented")
}
