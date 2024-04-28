package handler

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/auth/shared"
)

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
