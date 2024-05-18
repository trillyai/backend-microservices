package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// //////////////////////////////////////////////////////////////////////////////////
// CreateComment implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) CreateComment(ctx context.Context, req shared.CreateCommentRequest) (shared.CreateCommentResponse, error) {

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.CreateCommentResponse{}, errors.New("context not found")
	}

	req.Username = claims.UserName

	post, err := postgres.Read[tables.Post, tables.Post](ctx, map[string]interface{}{"Id": req.PostId})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.CreateCommentResponse{}, err
	}

	if post.Description == "" {
		return shared.CreateCommentResponse{}, errors.New("post not found")
	}

	resp, err := postgres.Create[shared.CreateCommentResponse, tables.Comment](ctx, req)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.CreateCommentResponse{}, err
	}

	return resp, nil

}

// //////////////////////////////////////////////////////////////////////////////////
// UpdateComment implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) UpdateComment(ctx context.Context, req shared.UpdateCommentRequest) (shared.UpdateCommentResponse, error) {

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.UpdateCommentResponse{}, errors.New("context not found")
	}

	comment, err := postgres.Read[tables.Comment, tables.Comment](ctx, map[string]interface{}{"Id": req.Id})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.UpdateCommentResponse{}, err
	}

	if comment.Username != claims.UserName {
		return shared.UpdateCommentResponse{}, errors.New("unauthorized to update")
	}

	resp, err := postgres.Update[shared.UpdateCommentResponse, tables.Comment](ctx, map[string]interface{}{"Id": req.Id}, req)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.UpdateCommentResponse{}, err
	}

	return resp, nil

}

// //////////////////////////////////////////////////////////////////////////////////
// DeleteComment implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) DeleteComment(ctx context.Context, req shared.DeleteCommentRequest) (shared.DeleteCommentReesponse, error) {

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.DeleteCommentReesponse{}, errors.New("context not found")
	}

	comment, err := postgres.Read[tables.Comment, tables.Comment](ctx, map[string]interface{}{"Id": req.Id})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.DeleteCommentReesponse{}, err
	}

	if comment.Username != claims.UserName {
		return shared.DeleteCommentReesponse{}, errors.New("unauthorized to delete")
	}

	resp, err := postgres.Delete[shared.DeleteCommentReesponse, tables.Comment](ctx, map[string]interface{}{"Id": req.Id})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.DeleteCommentReesponse{}, err
	}

	return resp, nil

}

// //////////////////////////////////////////////////////////////////////////////////
// GetComment implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) GetComment(ctx context.Context, commentId uuid.UUID) (shared.Comment, error) {

	comment, err := postgres.Read[shared.Comment, tables.Comment](ctx, map[string]interface{}{"Id": commentId})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.Comment{}, err
	}

	post, err := postgres.Read[tables.Post, tables.Post](ctx, map[string]interface{}{"Id": comment.PostId})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.Comment{}, err
	}

	if post.Id == uuid.Nil {
		return shared.Comment{}, errors.New("post associated with comment not found")
	}

	return comment, nil

}

// //////////////////////////////////////////////////////////////////////////////////
// GetComments implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) GetComments(ctx context.Context, uuidVal uuid.UUID, username string, forPostId bool, offset uint32, limit uint32) (shared.Comments, error) {

	var key string
	var value interface{}

	switch forPostId {
	case true:
		key = "PostId"
		value = uuidVal
	default:
		key = "Username"
		value = username
	}

	rule := map[string]interface{}{key: value}

	resp, err := postgres.PaginatedRead[[]shared.Comment, tables.Comment](ctx, rule, offset, limit)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.Comments{}, err
	}

	if key == "PostId" {
		post, err := postgres.Read[tables.Post, tables.Post](ctx, map[string]interface{}{"Id": uuidVal})
		if err != nil {
			r.logger.Error(err.Error())
			return shared.Comments{}, err
		}
		if post.Id == uuid.Nil {
			return shared.Comments{}, errors.New("post not found")
		}
	}

	ids, err := postgres.Read[[]utils.JustId, tables.Comment](ctx, rule)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.Comments{}, err
	}

	return shared.Comments{
		Comments:     resp,
		CommentCount: uint64(len(ids)),
	}, nil

}
