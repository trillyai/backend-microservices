package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/utils"
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
// @Summary GenerateFeed
// @Description GenerateFeed
// @Tags Feed
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param username query int false "Username"
// @Success 200 {object} shared.Feed
// @Failure 400 {string} string "Bad Request"
// @Router /feed [get]
func (handler handler) GenerateFeed(c *fiber.Ctx) error {

	offset, limit, err := utils.GetOffSetAndLimit(c.Query("offset"), c.Query("limit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	username := c.Query("username")

	resp, err := handler.svc.GenerateFeed(c.Context(), offset, limit, username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
