package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// CreatePost implements contracts.Handler.
// @Summary Create a new post
// @Description Create a new post with the given details
// @Tags posts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param post body shared.CreatePostRequest true "Post to create"
// @Success 200 {object} shared.CreatePostResponse
// @Failure 400 {string} string "Bad Request"
// @Router /posts [post]
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

// UpdatePost implements contracts.Handler.
// @Summary Update an existing post
// @Description Update an existing post with the given details
// @Tags posts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param post body shared.UpdatePostRequest true "Post to update"
// @Success 200 {object} shared.UpdatePostResponse
// @Failure 400 {string} string "Bad Request"
// @Router /posts [put]
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

// DeletePost implements contracts.Handler.
// @Summary Delete an existing post
// @Description Delete an existing post with the given details
// @Tags posts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param post body shared.DeletePostRequest true "Post to delete"
// @Success 200 {object} shared.DeletePostReesponse
// @Failure 400 {string} string "Bad Request"
// @Router /posts [delete]
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

// GetPost implements contracts.Handler.
// @Summary Get a post by ID
// @Description Get a post by its ID
// @Tags posts
// @Produce json
// @Param postId path string true "Post ID"
// @Success 200 {object} shared.Post
// @Failure 400 {string} string "Bad Request"
// @Router /posts/{postId} [get]
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

// GetPosts implements contracts.Handler.
// @Summary Get posts by user ID
// @Description Get a list of posts by user ID
// @Tags posts
// @Produce json
// @Param userId query string true "User ID"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} []shared.Post
// @Failure 400 {string} string "Bad Request"
// @Router /posts [get]
func (handler handler) GetPosts(c *fiber.Ctx) error {
	userIdStr := c.Query("userId")

	offset, limit, err := utils.GetOffSetAndLimit(c.Query("offset"), c.Query("limit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
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
