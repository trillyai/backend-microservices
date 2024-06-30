package handler

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/ping"
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

// Ping
// @Summary Check API status
// @Description Check if the API is running
// @Tags health
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (handler handler) Ping(c *fiber.Ctx) error {
	return ping.Ping(c)
}

// Register
// @Summary Register a new user
// @Description Register a new user with the system
// @Tags auth
// @Accept json
// @Produce json
// @Param user body shared.RegisterRequest true "User registration data"
// @Success 200 {object} shared.RegisterResponse
// @Router /register [post]
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

// Login
// @Summary Login a user
// @Description Login a user and return a token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body shared.LoginRequest true "User login data"
// @Success 200 {object} shared.LoginResponse
// @Router /login [post]
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

// Logout
// @Summary Logout a user
// @Description Logout a user and invalidate their token
// @Tags auth
// @Security ApiKeyAuth
// @Success 200 {object} shared.LogoutResponse
// @Router /logout [post]
func (handler handler) Logout(c *fiber.Ctx) error {

	resp, err := handler.svc.Logout(c.Context(), shared.LogoutRequest{})
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process logout: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}
