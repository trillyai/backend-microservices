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

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.UpdatePostResponse{}, errors.New("context not found")
	}

	post, err := postgres.Read[tables.Post, tables.Post](ctx, map[string]interface{}{"Id": req.Id, "Username": claims.UserName})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.UpdatePostResponse{}, err
	}

	if post.Id != req.Id {
		return shared.UpdatePostResponse{}, errors.New("post id does not match request id")
	}

	post.Description = req.Description

	resp, err := postgres.Update[shared.UpdatePostResponse, tables.Post](ctx, map[string]interface{}{"Id": req.Id}, post)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.UpdatePostResponse{}, err
	}

	return resp, err

}

// //////////////////////////////////////////////////////////////////////////////////
// DeletePost implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) DeletePost(ctx context.Context, req shared.DeletePostRequest) (shared.DeletePostReesponse, error) {

	claims := ctx.Value("user").(*auth.Claims)
	if claims.Name == "" {
		return shared.DeletePostReesponse{}, errors.New("context not found")
	}

	post, err := postgres.Read[tables.Post, tables.Post](ctx, map[string]interface{}{"Id": req.Id, "Username": claims.UserName})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.DeletePostReesponse{}, err
	}

	if post.Id != req.Id {
		return shared.DeletePostReesponse{}, errors.New("post id does not match request id")
	}

	resp, err := postgres.Delete[shared.DeletePostReesponse, tables.Post](ctx, map[string]interface{}{"Id": req.Id, "Username": claims.UserName})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.DeletePostReesponse{}, err
	}

	return resp, nil

}

// //////////////////////////////////////////////////////////////////////////////////
// GetPost implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) GetPost(ctx context.Context, postId uuid.UUID) (shared.Post, error) {

	resp, err := postgres.Read[shared.Post, tables.Post](ctx, map[string]interface{}{"Id": postId})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.Post{}, err
	}

	if resp.Id != postId {
		return shared.Post{}, errors.New("post not found")
	}

	return resp, nil

}

// //////////////////////////////////////////////////////////////////////////////////
// GetPosts implements contracts.Repository.
// //////////////////////////////////////////////////////////////////////////////////
func (r repository) GetPosts(ctx context.Context, userId uuid.UUID, offset uint32, limit uint32) ([]shared.Post, error) {

	user, err := postgres.Read[tables.User, tables.User](ctx, map[string]interface{}{"Id": userId})
	if err != nil {
		r.logger.Error(err.Error())
		return []shared.Post{}, err
	}

	if user.Username == "" {
		return []shared.Post{}, errors.New("user not exists")
	}

	resp, err := postgres.Read[[]shared.Post, tables.Post](ctx, map[string]interface{}{"Username": user.Username})
	if err != nil {
		r.logger.Error(err.Error())
		return []shared.Post{}, err
	}

	return resp, nil

}
