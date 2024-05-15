package handler

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/auth/contracts"
	"github.com/trillyai/backend-microservices/services/auth/shared"
)

type handler struct {
	svc    contracts.Service
	logger logger.Logger
}

func NewHandler(svc contracts.Service) contracts.Handler {
	return handler{
		svc:    svc,
		logger: *logger.NewLogger("auth-handler"),
	}
}

// //////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) Register(c *fiber.Ctx) error {

	var req shared.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	// Call the service's Register method to process the registration request.
	resp, err := handler.svc.Register(c.Context(), req)
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process registration: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Respond to the client with the registration response.
	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) Login(c *fiber.Ctx) error {

	var req shared.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.Login(c.Context(), req)
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process login: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) Logout(c *fiber.Ctx) error {

	resp, err := handler.svc.Logout(c.Context(), shared.LogoutRequest{})
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process logout: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}
