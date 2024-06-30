package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// CreateLike implements contracts.Handler.
// @Summary Create a new like
// @Description Create a new like for a post or a comment
// @Tags likes
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param like body shared.CreateLikeRequest true "Like to create"
// @Success 200 {object} shared.CreateLikeResponse
// @Failure 400 {string} string "Bad Request"
// @Router /likes [post]
func (handler handler) CreateLike(c *fiber.Ctx) error {
	type CreateLikeRequestWithStringDataTypes struct {
		PostId    string `json:"postId,omitempty"`
		CommentId string `json:"commentId,omitempty"`
	}

	var treq CreateLikeRequestWithStringDataTypes

	if err := c.BodyParser(&treq); err != nil {
		handler.logger.Error(fmt.Sprintf("failed to parse request body: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if treq.PostId == "" && treq.CommentId == "" {
		return c.Status(fiber.StatusBadRequest).SendString("postId or commentId required")
	}

	var req shared.CreateLikeRequest

	if treq.PostId != "" {
		id, err := uuid.Parse(treq.PostId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		req.PostId = id
	}

	if treq.CommentId != "" {
		id, err := uuid.Parse(treq.CommentId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		req.CommentId = id
	}

	resp, err := handler.svc.CreateLike(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

// DeleteLike implements contracts.Handler.
// @Summary Delete a like
// @Description Delete a like by providing the details
// @Tags likes
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param like body shared.DeleteLikeRequest true "Like to delete"
// @Success 200 {object} shared.DeleteLikeResponse
// @Failure 400 {string} string "Bad Request"
// @Router /likes [delete]
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

// GetLikes implements contracts.Handler.
// @Summary Get likes by post ID or comment ID
// @Description Get a list of likes by post ID or comment ID
// @Tags likes
// @Produce json
// @Param postId query string false "Post ID"
// @Param commentId query string false "Comment ID"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} shared.Likes
// @Failure 400 {string} string "Bad Request"
// @Router /likes [get]
func (handler handler) GetLikes(c *fiber.Ctx) error {
	postIdStr := c.Query("postId")
	commentIdStr := c.Query("commentId")

	offset, limit, err := utils.GetOffSetAndLimit(c.Query("offset"), c.Query("limit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if postIdStr == "" && commentIdStr == "" {
		return c.Status(fiber.StatusBadRequest).SendString("use postId or commentId to get likes")
	}

	if postIdStr != "" && commentIdStr != "" {
		return c.Status(fiber.StatusBadRequest).SendString("use just postId or commentId to get likes")
	}

	var forPostId bool = false
	var uid uuid.UUID

	if postIdStr != "" {
		postId, err := uuid.Parse(postIdStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		uid = postId
		forPostId = true
	}

	if commentIdStr != "" {
		commentId, err := uuid.Parse(commentIdStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		uid = commentId
		forPostId = false
	}

	resp, err := handler.svc.GetLikes(c.Context(), uid, forPostId, !forPostId, uint32(offset), uint32(limit))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
