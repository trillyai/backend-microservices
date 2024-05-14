package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// //////////////////////////////////////////////////////////////////////////////////
// CreateComment implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) CreateComment(ctx context.Context, req shared.CreateCommentRequest) (shared.CreateCommentResponse, error) {
	return s.repository.CreateComment(ctx, req)
}

// //////////////////////////////////////////////////////////////////////////////////
// UpdateComment implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) UpdateComment(ctx context.Context, req shared.UpdateCommentRequest) (shared.UpdateCommentResponse, error) {
	return s.repository.UpdateComment(ctx, req)
}

// //////////////////////////////////////////////////////////////////////////////////
// DeleteComment implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) DeleteComment(ctx context.Context, req shared.DeleteCommentRequest) (shared.DeleteCommentReesponse, error) {
	return s.repository.DeleteComment(ctx, req)
}

// //////////////////////////////////////////////////////////////////////////////////
// GetComment implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) GetComment(ctx context.Context, commentId uuid.UUID) (shared.Comment, error) {
	return s.repository.GetComment(ctx, commentId)
}

// //////////////////////////////////////////////////////////////////////////////////
// GetComments implements contracts.Service.
// //////////////////////////////////////////////////////////////////////////////////
func (s service) GetComments(ctx context.Context, uuid uuid.UUID, username string, forPostId bool, offset uint32, limit uint32) (shared.Comments, error) {
	return s.repository.GetComments(ctx, uuid, username, forPostId, offset, limit)
}
