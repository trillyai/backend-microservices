package shared

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateCommentRequest struct {
		PostId   uuid.UUID `json:"postId" validate:"required"`
		Comment  string    `json:"comment" validate:"required"`
		Username string    //extract from token
	}
	CreateCommentResponse struct {
		Id          uuid.UUID  `json:"id"`
		CreatedDate *time.Time `json:"createdDate"`
	}
)

type (
	UpdateCommentRequest struct {
		Id      uuid.UUID `json:"id" validate:"required"`
		Comment string    `json:"comment" validate:"required"`
	}
	UpdateCommentResponse struct {
		Id              uuid.UUID  `json:"id"`
		LastUpdatedDate *time.Time `json:"updateDate"`
	}
)

type (
	DeleteCommentRequest struct {
		Id uuid.UUID `json:"id"`
	}
	DeleteCommentResponse struct {
		DeletedDate *time.Time `json:"deletedDate"`
	}
)

type (
	Comment struct {
		Id       uuid.UUID `json:"id"`
		PostId   uuid.UUID `json:"postId"`
		Username string    `json:"userName"`
		Comment  string    `json:"comment"`
	}

	Comments struct {
		Comments     []Comment `json:"comments"`
		CommentCount uint64    `json:"commentCount"`
	}
)
