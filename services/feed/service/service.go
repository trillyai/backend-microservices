package service

import (
	"context"

	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/feed/contracts"
	"github.com/trillyai/backend-microservices/services/feed/shared"
)

type service struct {
	repository contracts.Repository
	logger     logger.Logger
}

func NewService(repository contracts.Repository) contracts.Service {
	return service{
		repository: repository,
		logger:     *logger.NewLogger("feed-service"),
	}
}

// GenerateFeed implements contracts.Service.
func (s service) GenerateFeed(ctx context.Context, offset uint32, limit uint32, username string) (shared.Feed, error) {
	return s.repository.GenerateFeed(ctx, offset, limit, username)
}
