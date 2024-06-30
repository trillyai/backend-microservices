package contracts

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GenerateFeed(c *fiber.Ctx) error
}
