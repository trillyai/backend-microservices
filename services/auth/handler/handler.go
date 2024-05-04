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
	// Parse the request body into a shared.RegisterRequest struct.
	var req shared.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}
	// Call the service's Register method to process the registration request.
	resp, err := handler.svc.Register(c.Context(), req)
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process registration: %v", err))
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	// Respond to the client with the registration response.
	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}

// //////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) Login(c *fiber.Ctx) error {
	var req shared.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	resp, err := handler.svc.Login(c.Context(), req)
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process login: %v", err))
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}

// //////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) Logout(c *fiber.Ctx) error {
	resp, err := handler.svc.Logout(c.Context(), shared.LogoutRequest{})
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process logout: %v", err))
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}
