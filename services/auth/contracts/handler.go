package contracts

import "github.com/gofiber/fiber"

type Handler interface {
	Register(c *fiber.Ctx)
}
