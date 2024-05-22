package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/trillyai/backend-microservices/core/bootstrap"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/core/env"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/middleware"
	"github.com/trillyai/backend-microservices/core/ping"
	"github.com/trillyai/backend-microservices/services/profile/contracts"
	"github.com/trillyai/backend-microservices/services/profile/handler"
	"github.com/trillyai/backend-microservices/services/profile/repository"
	"github.com/trillyai/backend-microservices/services/profile/service"

	// docs are generated by Swag CLI, you have to import them.
	_ "github.com/trillyai/backend-microservices/services/profile/docs"
)

const (
	profiles             = "/profiles"
	profilesWithUsername = profiles + "/:username"
)

func init() {
	bootstrap.SetUpEnvironment()
	if err := postgres.MigrateSchema(tables.User{}, tables.Session{}); err != nil {
		os.Exit(1)
	}
}

// @title Profile Server API
// @version 1.0
// @description This is the API documentation for the profile server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@profile-server.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8081
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
	logger := logger.NewLogger("profile-server")
	handler := getHandler()

	addMiddlewares(app)
	logger.Debug("Handler instance created")

	app.Get(ping.PingPath, ping.Ping)
	app.Get(profiles, handler.GetProfiles)
	app.Get(profilesWithUsername, handler.GetProfile)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	authMw := app.Group("", middleware.AuthMiddleware)
	authMw.Put(profiles, handler.UpdateProfile)
	authMw.Post(profiles, handler.UploadProfileImage)

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
