package service

import (
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/interaction/contracts"
)

type service struct {
	repository contracts.Repository
	logger     logger.Logger
}

func NewService(repository contracts.Repository) contracts.Service {
	return service{
		repository: repository,
		logger:     *logger.NewLogger("interaction-service"),
	}
}
