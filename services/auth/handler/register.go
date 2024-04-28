package handler

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/auth/shared"
)

// Register handles the registration request.
// @Summary Register a new user.
// @Description Register a new user with the provided details.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body shared.RegisterRequest true "Registration Request"
// @Success 200 {object} shared.RegisterResponse "Registration Response"
// @Router /register [post]
func (handler handler) Register(c *fiber.Ctx) error {
	// Parse the request body into a shared.RegisterRequest struct.
	var req shared.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest)
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		c.JSON(err.Error())
		return err
	}

	if err := utils.ValidateStruct(req); err != nil {
		c.Status(fiber.StatusBadRequest)
		handler.logger.Error(err.Error())
		c.JSON(err.Error())
		return err
	}

	// Call the service's Register method to process the registration request.
	resp, err := handler.svc.Register(c.Context(), req)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		handler.logger.Error(fmt.Sprintf("failed to process registration: %v", err))
		c.JSON(err.Error())
		return err
	}

	// Respond to the client with the registration response.
	c.Status(fiber.StatusOK)
	c.JSON(resp)
	return nil
}
