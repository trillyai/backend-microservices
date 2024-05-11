package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	if err := utils.ValidateStruct(req); err != nil {
		handler.logger.Error(err.Error())
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	resp, err := handler.svc.CreatePost(c.Context(), req)
	if err != nil {
		handler.logger.Error(fmt.Sprintf("failed to process create-post: %v", err))
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil

}

// //////////////////////////////////////////////////////////////////////////////////
// UpdatePost implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) UpdatePost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// DeletePost implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) DeletePost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// GetPost implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) GetPost(c *fiber.Ctx) error {
	panic("unimplemented")
}

// //////////////////////////////////////////////////////////////////////////////////
// GetPosts implements contracts.Handler.
// //////////////////////////////////////////////////////////////////////////////////
func (handler handler) GetPosts(c *fiber.Ctx) error {
	panic("unimplemented")
}
