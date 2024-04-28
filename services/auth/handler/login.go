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

	resp, err := handler.svc.Login(c.Context(), req)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		handler.logger.Error(fmt.Sprintf("failed to process login: %v", err))
		c.JSON(err.Error())
		return err
	}

	c.Status(fiber.StatusOK)
	c.JSON(resp)
	return nil
}
