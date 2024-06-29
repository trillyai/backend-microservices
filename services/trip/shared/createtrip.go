package shared

import "github.com/google/uuid"

type (
	CreateTripRequest struct {
		Areas    string `validate:"required"`
		Filters  string `validate:"required"`
		Distance string
	}

	CreateTripResponse struct {
		Id uuid.UUID `json:"id"`
		Root
	}
)
