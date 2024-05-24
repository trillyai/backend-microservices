package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// CreateComment implements contracts.Handler.
// @Summary Create a new comment
// @Description Create a new comment with the given details
// @Tags comments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param comment body shared.CreateCommentRequest true "Comment to create"
// @Success 200 {object} shared.CreateCommentResponse
// @Failure 400 {string} string "Bad Request"
// @Router /comments [post]
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

// UpdateComment implements contracts.Handler.
// @Summary Update an existing comment
// @Description Update an existing comment with the given details
// @Tags comments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param comment body shared.UpdateCommentRequest true "Comment to update"
// @Success 200 {object} shared.UpdateCommentResponse
// @Failure 400 {string} string "Bad Request"
// @Router /comments [put]
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

// DeleteComment implements contracts.Handler.
// @Summary Delete an existing comment
// @Description Delete an existing comment with the given details
// @Tags comments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param comment body shared.DeleteCommentRequest true "Comment to delete"
// @Success 200 {object} shared.DeleteCommentResponse
// @Failure 400 {string} string "Bad Request"
// @Router /comments [delete]
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

// GetComment implements contracts.Handler.
// @Summary Get a comment by ID
// @Description Get a comment by its ID
// @Tags comments
// @Produce json
// @Param commentId path string true "Comment ID"
// @Success 200 {object} shared.Comment
// @Failure 400 {string} string "Bad Request"
// @Router /comments/{commentId} [get]
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

// GetComments implements contracts.Handler.
// @Summary Get comments by username or post ID
// @Description Get a list of comments by username or post ID
// @Tags comments
// @Produce json
// @Param username query string false "Username"
// @Param postId query string false "Post ID"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} shared.Comments
// @Failure 400 {string} string "Bad Request"
// @Router /comments [get]
func (handler handler) GetComments(c *fiber.Ctx) error {
	username := c.Query("username")
	postIdStr := c.Query("postId")

	offset, limit, err := utils.GetOffSetAndLimit(c.Query("offset"), c.Query("limit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
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
