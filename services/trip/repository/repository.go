package repository

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/trillyai/backend-microservices/core/auth"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
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
	areas := strings.Split(req.Areas, ",")
	filters := strings.Split(req.Filters, ",")

	output, err := sendRequest(areas, filters, req.Distance)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.CreateTripResponse{}, err
	}

	if len(output.Data) == 0 {
		return shared.CreateTripResponse{}, errors.New("no trip for specified filters and areas")
	}

	jsonRaw, _ := structToJSONRawMessage(output)
	rowTrip, err := postgres.Create[tables.Trip, tables.Trip](ctx, tables.Trip{DataJson: jsonRaw})
	if err != nil {
		r.logger.Error(err.Error())
		return shared.CreateTripResponse{}, err
	}

	for _, filter := range filters {
		postgres.Create[struct{}, tables.TripInterest](ctx, tables.TripInterest{
			TripId:   rowTrip.Id,
			Interest: filter,
		})
	}

	return shared.CreateTripResponse{Root: output}, nil
}

// StructToJSONRawMessage converts a struct to json.RawMessage
func structToJSONRawMessage(v interface{}) (json.RawMessage, error) {
	// Marshal the struct to JSON
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	// Convert the JSON bytes to json.RawMessage
	return json.RawMessage(data), nil
}
