package contracts

import (
	"context"

	"github.com/trillyai/backend-microservices/services/trip/shared"
)

type Repository interface {
	CreateTrip(ctx context.Context, req shared.CreateTripRequest) (shared.CreateTripResponse, error)
}