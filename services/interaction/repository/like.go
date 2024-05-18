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
// CreateLike implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) CreateLike(ctx context.Context, req shared.CreateLikeRequest) (shared.CreateLikeResponse, error) {

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.CreateLikeResponse{}, errors.New("context not found")
	}

	req.Username = claims.UserName

	// post like request
	if req.PostId != uuid.Nil {
		post, err := postgres.Read[tables.Post, tables.Post](ctx, map[string]interface{}{"Id": req.PostId})
		if err != nil {
			r.logger.Error(err.Error())
			return shared.CreateLikeResponse{}, err
		}

		if post.Id == uuid.Nil {
			return shared.CreateLikeResponse{}, errors.New("post not found")
		}

		isLiked, err := postgres.Read[tables.Like, tables.Like](ctx, map[string]interface{}{"PostId": req.PostId})
		if err != nil {
			r.logger.Error(err.Error())
			return shared.CreateLikeResponse{}, err
		}

		if isLiked.Id != uuid.Nil {
			return shared.CreateLikeResponse{}, errors.New("already liked")
		}

		resp, err := postgres.Create[shared.CreateLikeResponse, tables.Like](ctx, req)
		if err != nil {
			r.logger.Error(err.Error())
			return shared.CreateLikeResponse{}, err
		}
		return resp, nil
	}

	// comment like request
	comment, err := postgres.Read[tables.Comment, tables.Comment](ctx, map[string]interface{}{"Id": req.CommentId})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.CreateLikeResponse{}, err
	}

	if comment.Id == uuid.Nil {
		return shared.CreateLikeResponse{}, errors.New("comment not found")
	}

	isLiked, err := postgres.Read[tables.Like, tables.Like](ctx, map[string]interface{}{"CommentId": req.CommentId})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.CreateLikeResponse{}, err
	}

	if isLiked.Id != uuid.Nil {
		return shared.CreateLikeResponse{}, errors.New("already liked")
	}

	resp, err := postgres.Create[shared.CreateLikeResponse, tables.Like](ctx, req)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.CreateLikeResponse{}, err
	}
	return resp, nil

}

// //////////////////////////////////////////////////////////////////////////////////
// DeleteLike implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) DeleteLike(ctx context.Context, req shared.DeleteLikeRequest) (shared.DeleteLikeResponse, error) {

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.DeleteLikeResponse{}, errors.New("context not found")
	}

	like, err := postgres.Read[tables.Like, tables.Like](ctx, map[string]interface{}{"Id": req.Id})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.DeleteLikeResponse{}, err
	}

	if like.Username != claims.UserName {
		return shared.DeleteLikeResponse{}, errors.New("unauthorized to delete")
	}

	resp, err := postgres.Delete[shared.DeleteLikeResponse, tables.Like](ctx, map[string]interface{}{"Id": req.Id})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.DeleteLikeResponse{}, err
	}

	return resp, nil

}

// //////////////////////////////////////////////////////////////////////////////////
// GetLikes implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) GetLikes(ctx context.Context, uuid uuid.UUID, forPostId bool, forCommentId bool, offset uint32, limit uint32) ([]shared.Like, error) {

	var key string

	switch forPostId {
	case true:
		key = "PostId"
	default:
		key = "CommentId"
	}

	resp, err := postgres.PaginatedRead[[]shared.Like, tables.Like](ctx, map[string]interface{}{key: uuid}, offset, limit)
	if err != nil {
		r.logger.Error(err.Error())
		return []shared.Like{}, err
	}

	return resp, nil

}
