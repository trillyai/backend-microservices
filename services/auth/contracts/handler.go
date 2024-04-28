package contracts

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
