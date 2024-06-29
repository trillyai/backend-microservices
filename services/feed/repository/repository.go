package repository

import (
	"context"

	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/feed/contracts"
	"github.com/trillyai/backend-microservices/services/feed/shared"
)

type repository struct {
	logger logger.Logger
}

func NewRepository() contracts.Repository {
	return repository{
		logger: *logger.NewLogger("feed-repository"),
	}
}

// GenerateFeed implements contracts.Repository.
func (r repository) GenerateFeed(ctx context.Context, offset uint32, limit uint32) (shared.Feed, error) {
	panic("unimplemented")
}
