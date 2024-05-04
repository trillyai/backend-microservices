package ping

import "github.com/gofiber/fiber/v2"

const PingPath = "/ping"

func Ping(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK).Send([]byte("pong dude"))
	return nil
}
