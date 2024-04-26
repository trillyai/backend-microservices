package server

import (
	"github.com/gofiber/fiber"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/services/auth/handler"
	"github.com/trillyai/backend-microservices/services/auth/repository"
	"github.com/trillyai/backend-microservices/services/auth/service"
)

func GetServerApp() *fiber.App {
	app := fiber.New()
	logger := logger.NewLogger("auth-server")

	logger.Info("creating server app...")

	repo := repository.NewRepository()
	logger.Debug("repository instance created")

	service := service.NewService(repo)
	logger.Debug("service instance created")

	handler := handler.NewHandler(service)
	logger.Debug("Handler instance created")

	app.Post("/register", handler.Register)

	app.Get("/ping", func(c *fiber.Ctx) {
		c.Write("pong")
		c.Status(fiber.StatusOK)
	})
	logger.Info("Server app initialization completed.")

	return app
}
