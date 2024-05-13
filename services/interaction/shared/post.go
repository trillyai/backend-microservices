package shared

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreatePostRequest struct {
		TripId      uuid.UUID `json:"tripId" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Username    string
	}
	CreatePostResponse struct {
		Id          uuid.UUID  `json:"id"`
		CreatedDate *time.Time `json:"createdDate"`
	}
)

type (
	UpdatePostRequest struct {
		Id          uuid.UUID `json:"id" validate:"required"`
		TripId      uuid.UUID `json:"tripId" validate:"required"`
		Description string    `json:"description" validate:"required"`
	}
	UpdatePostResponse struct {
		Id          uuid.UUID  `json:"id"`
		TripId      uuid.UUID  `json:"tripId"`
		Description string     `json:"description"`
		UpdatedDate *time.Time `json:"updateDate"`
	}
)

type (
	DeletePostRequest struct {
		Id uuid.UUID `json:"id" validate:"required"`
	}
	DeletePostReesponse struct {
		DeletedDate *time.Time `json:"deletedDate"`
	}
)

type Post struct {
	Id          uuid.UUID `json:"id"`
	TripId      uuid.UUID `json:"tripId"`
	Username    string    `json:"userName"`
	Description string    `json:"description"`
}
