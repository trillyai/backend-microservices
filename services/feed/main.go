package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/trillyai/backend-microservices/core/bootstrap"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/env"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/ping"
	"github.com/trillyai/backend-microservices/services/feed/contracts"
	"github.com/trillyai/backend-microservices/services/feed/handler"
	"github.com/trillyai/backend-microservices/services/feed/repository"
	"github.com/trillyai/backend-microservices/services/feed/service"
)

const (
	feed = "/feed"

	swag = "/swagger/*"
)

func init() {
	bootstrap.SetUpEnvironment()
	if err := postgres.MigrateSchema(); err != nil {
		os.Exit(1)
	}
}

// @title Trip Server API
// @version 1.0
// @description This is the API documentation for the feed server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@feed-server.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8084
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
		if err := app.Listen(env.HttpPort); err != nil && err != http.ErrServerClosed {
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
	logger := logger.NewLogger("feed-server")
	handler := getHandler()

	addMiddlewares(app)
	logger.Debug("Handler instance created")

	app.Get(ping.PingPath, ping.Ping)
	app.Get(swag, swagger.HandlerDefault) // default
	app.Get(feed, handler.GenerateFeed)

	logger.Info("Server app initialization completed.")
	return app
}

func addMiddlewares(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Content-Type, Authorization",
	}))
	app.Use(fiberLogger.New())
	app.Use(requestid.New())
}

func getHandler() contracts.Handler {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	return handler
}
