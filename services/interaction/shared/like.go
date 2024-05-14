package shared

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateLikeRequest struct {
		PostId    uuid.UUID `json:"postId,omitempty"`
		CommentId uuid.UUID `json:"commentId,omitempty"`
		Username  string
	}
	CreateLikeResponse struct {
		Id          uuid.UUID  `json:"id"`
		CreatedDate *time.Time `json:"createdDate"`
	}
)

type (
	DeleteLikeRequest struct {
		Id uuid.UUID `json:"id" validate:"required"`
	}
	DeleteLikeResponse struct {
		Id          uuid.UUID  `json:"id"`
		DeletedDate *time.Time `json:"deletedDate"`
	}
)

type Like struct {
	Id        uuid.UUID `json:"id"`
	PostId    uuid.UUID `json:"postId,omitempty"`
	CommentId uuid.UUID `json:"commentId,omitempty"`
	Username  string    `json:"userName"`
}
