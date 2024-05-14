package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// //////////////////////////////////////////////////////////////////////////////////
// CreateLike implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) CreateLike(c *fiber.Ctx) error {

	var req shared.CreateLikeRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if req.PostId == req.CommentId && req.PostId == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).SendString("postId or commentId required")
	}

	resp, err := handler.svc.CreateLike(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// DeleteLike implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) DeleteLike(c *fiber.Ctx) error {

	var req shared.DeleteLikeRequest

	if err := c.BodyParser(&req); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	resp, err := handler.svc.DeleteLike(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}

// //////////////////////////////////////////////////////////////////////////////////
// GetLikes implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) GetLikes(c *fiber.Ctx) error {

	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	postIdStr := c.Query("postId")
	commentIdStr := c.Query("commentId")

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
	commentId, _ := uuid.Parse(commentIdStr)

	if postId != uuid.Nil && commentId != uuid.Nil {
		return c.Status(fiber.StatusBadRequest).SendString("use just postId or commentId")
	}

	var forPostId bool = false
	var uid uuid.UUID

	if postId != uuid.Nil {
		forPostId = true
		uid = postId
	} else {
		forPostId = false
		uid = commentId
	}

	resp, err := handler.svc.GetLikes(c.Context(), uid, forPostId, !forPostId, uint32(offset), uint32(limit))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)

}
