package repository

import (
	"context"
	"errors"

	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/trip/contracts"
	"github.com/trillyai/backend-microservices/services/trip/shared"
)

type repository struct {
	logger logger.Logger
}

func NewRepository() contracts.Repository {
	return repository{
		logger: *logger.NewLogger("trip-repository"),
	}
}

// CreateTrip implements contracts.Repository.
func (r repository) CreateTrip(ctx context.Context, req shared.CreateTripRequest) (shared.CreateTripResponse, error) {
	claims := ctx.Value("user").(*auth.Claims)
	if claims.UserName == "" {
		return shared.CreateTripResponse{}, errors.New("context not found")
	}

	return shared.CreateTripResponse{}, nil
}
