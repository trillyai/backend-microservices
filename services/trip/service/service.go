package service

import (
	"context"

	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/trip/contracts"
	"github.com/trillyai/backend-microservices/services/trip/shared"
)

type service struct {
	repository contracts.Repository
	logger     logger.Logger
}

func NewService(repository contracts.Repository) contracts.Service {
	return service{
		repository: repository,
		logger:     *logger.NewLogger("trip-service"),
	}
}

// CreateTrip implements contracts.Service.
func (s service) CreateTrip(ctx context.Context, req shared.CreateTripRequest) (shared.CreateTripResponse, error) {
	return s.repository.CreateTrip(ctx, req)
}
