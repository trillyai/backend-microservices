package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/ping"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/trip/contracts"
	"github.com/trillyai/backend-microservices/services/trip/shared"
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

// CreateTrip
// @Summary Creates Trip
// @Description creates trip
// @Tags trip
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body shared.CreateTripRequest true "Create Trip Request"
// @Success 200 {object} shared.CreateTripResponse
// @Failure 400 {string} string "Bad Request"
// @Router /trip [post]
func (handler handler) CreateTrip(c *fiber.Ctx) error {
	var req shared.CreateTripRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.CreateTrip(c.Context(), req)
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process create-trip %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}
