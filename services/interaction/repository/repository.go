package repository

import (
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/interaction/contracts"
)

type repository struct {
	logger logger.Logger
}

func NewRepository() contracts.Repository {
	return repository{
		logger: *logger.NewLogger("interaction-repository"),
	}
}
