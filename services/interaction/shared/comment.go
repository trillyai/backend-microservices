package shared

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateCommentRequest struct {
		PostId  uuid.UUID `json:"postId" validate:"required"`
		Comment string    `json:"comment" validate:"required"`
		// Username string    `json:"userName" validate:"required"` //extract from token
	}
	CreateCommentResponse struct {
		Id          uuid.UUID  `json:"id"`
		CreatedDate *time.Time `json:"createdDate"`
	}
)

type (
	UpdateCommentRequest struct {
		TripId      uuid.UUID `json:"tripId" validate:"required"`
		Description string    `json:"description" validate:"required"`
		// Username    string    `json:"userName" validate:"required"` //extract from token
	}
	UpdateCommentResponse struct {
		Id          uuid.UUID  `json:"id"`
		UpdatedDate *time.Time `json:"updateDate"`
	}
)

type (
	DeleteCommentRequest struct {
		Id uuid.UUID `json:"id"`
	}
	DeleteCommentReesponse struct {
		Id          uuid.UUID  `json:"id"`
		DeletedDate *time.Time `json:"deletedDate"`
	}
)

type Comment struct {
	Id          uuid.UUID `json:"id"`
	TripId      uuid.UUID `json:"tripId"`
	Username    string    `json:"userName"`
	Description string    `json:"description"`
}
