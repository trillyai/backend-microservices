package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/interaction/shared"
)

// //////////////////////////////////////////////////////////////////////////////////
// CreateLike implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
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

	postIdStr := c.Query("postId")
	commentIdStr := c.Query("commentId")

	offset, limit, err := utils.GetOffSetAndLimit(c.Query("offset"), c.Query("limit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
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
