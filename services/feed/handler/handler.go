package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/feed/contracts"
)

type handler struct {
	svc    contracts.Service
	logger logger.Logger
}

func NewHandler(svc contracts.Service) contracts.Handler {
	return handler{
		svc:    svc,
		logger: *logger.NewLogger("feed-handler"),
	}
}

// GenerateFeed implements contracts.Handler.
func (h handler) GenerateFeed(c *fiber.Ctx) error {
	panic("unimplemented")
}
