package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// //////////////////////////////////////////////////////////////////////////////////
// CreatePost implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) CreatePost(ctx context.Context, req shared.CreatePostRequest) (shared.CreatePostResponse, error) {

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.CreatePostResponse{}, errors.New("context not found")
	}

	req.Username = claims.UserName

	resp, err := postgres.Create[shared.CreatePostResponse, tables.Post](ctx, req)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.CreatePostResponse{}, err
	}

	return resp, nil

}

// //////////////////////////////////////////////////////////////////////////////////
// UpdatePost implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) UpdatePost(ctx context.Context, req shared.UpdatePostRequest) (shared.UpdatePostResponse, error) {
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// DeletePost implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) DeletePost(ctx context.Context, req shared.DeletePostRequest) (shared.DeletePostReesponse, error) {
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// GetPost implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) GetPost(ctx context.Context, postId uuid.UUID) (shared.Post, error) {
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// GetPosts implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) GetPosts(ctx context.Context, userId uuid.UUID, offset uint32, limit uint32) ([]shared.Post, error) {
	panic("unimplemented")
}
