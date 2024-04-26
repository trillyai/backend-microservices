package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gofiber/fiber"
	"github.com/trillyai/backend-microservices/core/bootstrap"
	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/services/auth/server"
)

func init() {
	bootstrap.SetUpEnvironment()
	postgres.MigrateSchema(tables.User{})
}

func main() {
	app := server.GetServerApp()
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
