package contracts

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetProfiles(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
	UploadProfileImage(c *fiber.Ctx) error

	CreateUserInterest(c *fiber.Ctx) error
	DeleteUserInterest(c *fiber.Ctx) error
	GetUserInterests(c *fiber.Ctx) error
	GetIntersts(c *fiber.Ctx) error

	Ping(c *fiber.Ctx) error
}
