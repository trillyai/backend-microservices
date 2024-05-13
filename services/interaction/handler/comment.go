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
// CreateComment implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) CreateComment(c *fiber.Ctx) error {

	var req shared.CreateCommentRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.CreateComment(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// UpdateComment implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) UpdateComment(c *fiber.Ctx) error {

	var req shared.UpdateCommentRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.UpdateComment(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// DeleteComment implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) DeleteComment(c *fiber.Ctx) error {

	var req shared.DeleteCommentRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.DeleteComment(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// GetComment implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) GetComment(c *fiber.Ctx) error {

	commentIdStr := c.Params("commentId")

	if strings.TrimSpace(commentIdStr) == "" {
		handler.logger.Error("commentId param required")
		return c.Status(fiber.StatusBadRequest).SendString("commentId param required")
	}

	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.GetComment(c.Context(), commentId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// GetComments implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) GetComments(c *fiber.Ctx) error {

	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	username := c.Query("username")
	postIdStr := c.Query("postId")

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

	postId, _ := uuid.Parse(postIdStr)

	if strings.TrimSpace(username) == "" && postId == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).SendString("postId or username required")
	}

	var forPostId bool = true
	if strings.TrimSpace(username) != "" {
		forPostId = false
	}

	resp, err := handler.svc.GetComments(c.Context(), postId, username, forPostId, uint32(offset), uint32(limit))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}
