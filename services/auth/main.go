package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/trillyai/backend-microservices/core/bootstrap"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/middleware"
	"github.com/trillyai/backend-microservices/services/auth/contracts"
	"github.com/trillyai/backend-microservices/services/auth/handler"
	"github.com/trillyai/backend-microservices/services/auth/repository"
	"github.com/trillyai/backend-microservices/services/auth/service"
)

func init() {
	bootstrap.SetUpEnvironment()
	if err := postgres.MigrateSchema(tables.User{}, tables.Session{}); err != nil {
		os.Exit(1)
	}
}

func main() {
	app := GetServerApp()
	if err := StartServerWithGracefulShutdown(app); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func StartServerWithGracefulShutdown(app *fiber.App) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := app.Listen(bootstrap.Configs["HTTP_PORT"]); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v\n", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		return err
	}

	log.Println("Server shutdown completed")
	return nil
}

func GetServerApp() *fiber.App {
	app := fiber.New()
	logger := logger.NewLogger("auth-server")

	handler := getHandler()
	logger.Debug("Handler instance created")

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	app.Get("/ping", func(c *fiber.Ctx) error {
		c.Write([]byte("pong dude"))
		c.Status(fiber.StatusOK)
		return nil
	})

	authApp := app.Group("", middleware.AuthMiddleware)
	authApp.Get("/get-profile", handler.GetProfile)

	logger.Info("Server app initialization completed.")
	return app
}

func getHandler() contracts.Handler {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	return handler
}
