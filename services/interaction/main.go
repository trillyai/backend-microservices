package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/trillyai/backend-microservices/core/bootstrap"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/core/env"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/middleware"
	"github.com/trillyai/backend-microservices/core/ping"
	"github.com/trillyai/backend-microservices/services/interaction/contracts"
	"github.com/trillyai/backend-microservices/services/interaction/handler"
	"github.com/trillyai/backend-microservices/services/interaction/repository"
	"github.com/trillyai/backend-microservices/services/interaction/service"
)

const (
	posts  = "/posts"
	postId = "/:postId"

	comments = "/comments"
	likes    = "/likes"
)

func init() {
	bootstrap.SetUpEnvironment()
	if err := postgres.MigrateSchema(tables.Post{}, tables.Comment{}, tables.Like{}); err != nil {
		os.Exit(1)
	}
}

func main() {
	if err := StartServerWithGracefulShutdown(GetServerApp()); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func GetServerApp() *fiber.App {
	app := fiber.New()
	logger := logger.NewLogger("interaction-server")
	handler := getHandler()
	addMiddlewares(app)
	logger.Debug("Handler instance created")

	app.Get(ping.PingPath, ping.Ping)

	authApp := app.Group("", middleware.AuthMiddleware)
	servePostEndpoints(app, authApp, handler)

	logger.Info("Server app initialization completed.")
	return app
}

func servePostEndpoints(app *fiber.App, authApp fiber.Router, handler contracts.Handler) {
	app.Get(posts, handler.GetPosts)
	app.Get(posts+postId, handler.GetPost)

	authApp.Post(posts, handler.CreatePost)
	authApp.Put(posts, handler.UpdatePost)
	authApp.Delete(posts, handler.DeletePost)
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

func addMiddlewares(app *fiber.App) {
	app.Use(cors.New())
	app.Use(fiberLogger.New())
	app.Use(requestid.New())
}

func getHandler() contracts.Handler {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	return handler
}
