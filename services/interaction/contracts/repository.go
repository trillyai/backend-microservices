package contracts

import (
	"context"

	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

type Repository interface {
	postRepository
	commentRepository
	likeRepository
}

type postRepository interface {
	CreatePost(ctx context.Context, req shared.CreatePostRequest) (shared.CreatePostResponse, error)
	UpdatePost(ctx context.Context, req shared.UpdatePostRequest) (shared.UpdatePostResponse, error)
	DeletePost(ctx context.Context, req shared.DeletePostRequest) (shared.DeletePostReesponse, error)
	GetPost(ctx context.Context, postId uuid.UUID) (shared.Post, error)
	GetPosts(ctx context.Context, userId uuid.UUID, offset, limit uint32) ([]shared.Post, error)
}

type commentRepository interface {
	CreateComment(ctx context.Context, req shared.CreateCommentRequest) (shared.CreateCommentResponse, error)
	UpdateComment(ctx context.Context, req shared.UpdateCommentRequest) (shared.UpdateCommentResponse, error)
	DeleteComment(ctx context.Context, req shared.DeleteCommentRequest) (shared.DeleteCommentReesponse, error)
	GetComment(ctx context.Context, commentId uuid.UUID) (shared.Comment, error)
	GetComments(ctx context.Context, uuid uuid.UUID, forPostId, forUserId bool, offset, limit uint32) ([]shared.Comment, error)
}

type likeRepository interface {
	CreateLike(ctx context.Context, req shared.CreateLikeRequest) (shared.CreateLikeResponse, error)
	DeleteLike(ctx context.Context, req shared.DeleteLikeRequest) (shared.DeleteLikeReesponse, error)
	GetLikes(ctx context.Context, uuid uuid.UUID, forPostId, forUserId bool, offset, limit uint32) (shared.CreateLikeResponse, error)
}
