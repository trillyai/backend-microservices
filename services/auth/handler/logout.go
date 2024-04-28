package handler

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/services/auth/shared"
)

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
