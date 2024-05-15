package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/profile/contracts"
	"github.com/trillyai/backend-microservices/services/profile/shared"
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

// //////////////////////////////////////////////////////////////////////////////////
// GetProfile implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) GetProfile(c *fiber.Ctx) error {

	username := c.Params("username")
	if strings.TrimSpace(username) == "" {
		return c.Status(fiber.StatusBadRequest).SendString("username required")
	}

	resp, err := handler.svc.GetProfileByUsername(c.Context(), username)
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process get profile: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())

	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// UploadProfileImage implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) UploadProfileImage(c *fiber.Ctx) error {
	return nil
}

// //////////////////////////////////////////////////////////////////////////////////
// GetProfiles implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) GetProfiles(c *fiber.Ctx) error {

	offset, limit, err := utils.GetOffSetAndLimit(c.Query("offset"), c.Query("limit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.GetProfiles(c.Context(), uint32(offset), uint32(limit))
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process get profiles: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// UpdateProfile implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) UpdateProfile(c *fiber.Ctx) error {

	var req shared.UpdateProfileRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.UpdateProfile(c.Context(), req)
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process update-profile: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}
