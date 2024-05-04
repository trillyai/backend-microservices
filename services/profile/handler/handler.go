package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/profile/contracts"
)

type handler struct {
	svc    contracts.Service
	logger logger.Logger
}

func NewHandler(svc contracts.Service) contracts.Handler {
	return handler{
		svc:    svc,
		logger: *logger.NewLogger("profile-handler"),
	}
}

// GetProfile implements contracts.Handler.
func (h handler) GetProfile(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetProfiles implements contracts.Handler.
func (h handler) GetProfiles(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateProfile implements contracts.Handler.
func (h handler) UpdateProfile(c *fiber.Ctx) error {
	panic("unimplemented")
}
