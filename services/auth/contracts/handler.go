package contracts

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Ping(c *fiber.Ctx) error
}
