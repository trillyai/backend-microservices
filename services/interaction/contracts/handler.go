package contracts

import "github.com/gofiber/fiber/v2"

type Handler interface {
	postHandler
	commentHandler
	likeHandler
}

type postHandler interface {
	CreatePost(c *fiber.Ctx) error
	UpdatePost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error
	GetPost(fiber.Ctx) error  // postId
	GetPosts(fiber.Ctx) error // userId
}

type commentHandler interface {
	CreateComment(c *fiber.Ctx) error
	UpdateComment(c *fiber.Ctx) error
	DeleteComment(c *fiber.Ctx) error
	GetComment(c *fiber.Ctx) error  // commentId
	GetComments(c *fiber.Ctx) error // userId, postId
}

type likeHandler interface {
	CreateLike(c *fiber.Ctx) error
	DeleteLike(c *fiber.Ctx) error
	GetLikes(c *fiber.Ctx) error // postId, userId
}
