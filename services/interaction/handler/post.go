package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// //////////////////////////////////////////////////////////////////////////////////
// CreatePost implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) CreatePost(c *fiber.Ctx) error {

	var req shared.CreatePostRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.CreatePost(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// UpdatePost implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) UpdatePost(c *fiber.Ctx) error {

	var req shared.UpdatePostRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.UpdatePost(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// DeletePost implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) DeletePost(c *fiber.Ctx) error {

	var req shared.DeletePostRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.DeletePost(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// GetPost implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) GetPost(c *fiber.Ctx) error {

	postIdStr := c.Params("postId")

	if strings.TrimSpace(postIdStr) == "" {
		handler.logger.Error("postId param required")
		return c.Status(fiber.StatusBadRequest).SendString("postId param required")
	}

	postId, err := uuid.Parse(postIdStr)
	if err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.GetPost(c.Context(), postId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// GetPosts implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) GetPosts(c *fiber.Ctx) error {

	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	userIdStr := c.Query("userId")

	offset, err := strconv.ParseUint(offsetStr, 10, 32)
	if err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Invalid offset value")
	}

	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Invalid limit value")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.GetPosts(c.Context(), userId, uint32(offset), uint32(limit))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}
