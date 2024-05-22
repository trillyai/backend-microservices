package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/ping"
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

// API Ping
// @Summary Check API status
// @Description Check if the API is running
// @Tags health
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (handler handler) Ping(c *fiber.Ctx) error {
	return ping.Ping(c)
}

// GetProfile
// @Summary Get profile by username
// @Description Get profile information by username
// @Tags profiles
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} shared.GetProfileResponse
// @Router /profiles/{username} [get]
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

// GetProfiles
// @Summary Get profiles with pagination
// @Description Get a list of profiles with pagination support
// @Tags profiles
// @Accept json
// @Produce json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {object} []shared.GetProfileResponse
// @Router /profiles [get]
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

// UpdateProfile
// @Summary Update profile
// @Description Update profile information
// @Tags profiles
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body shared.UpdateProfileRequest true "Update Profile Request"
// @Success 200 {object} shared.UpdateProfileResponse
// @Router /profiles [put]
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

// UploadProfileImage
// @Summary Upload profile image
// @Description Upload a profile image
// @Tags profiles
// @Security ApiKeyAuth
// @Accept mpfd
// @Produce json
// @Param file formData file true "Profile Image File"
// @Success 200 {object} shared.UploadProfileImageResponse
// @Router /profiles/image [post]
func (handler handler) UploadProfileImage(c *fiber.Ctx) error {

	var req shared.UploadProfileImageRequest

	file, err := c.FormFile("file")
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse form file: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	req.File = file

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.UploadProfileImage(c.Context(), req)
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process update-profile: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}
