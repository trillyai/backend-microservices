package contracts

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateTrip(c *fiber.Ctx) error
	Ping(c *fiber.Ctx) error
}
