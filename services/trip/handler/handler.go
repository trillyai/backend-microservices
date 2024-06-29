package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/ping"
	"github.com/trillyai/backend-microservices/services/trip/contracts"
)

type handler struct {
	svc    contracts.Service
	logger logger.Logger
}

func NewHandler(svc contracts.Service) contracts.Handler {
	return handler{
		svc:    svc,
		logger: *logger.NewLogger("trip-handler"),
	}
}

// API Ping
// @Summary Check API status
// @Description Check if the API is running
// @Tags health
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (handler handler) Ping(c *fiber.Ctx) error {
	return ping.Ping(c)
}

// CreateTrip implements contracts.Handler.
func (handler handler) CreateTrip(c *fiber.Ctx) error {
	panic("unimplemented")
}
