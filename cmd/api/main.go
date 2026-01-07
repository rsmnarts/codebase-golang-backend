package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rsmnarts/codebase-golang-backend/internal/delivery/http"
	"github.com/rsmnarts/codebase-golang-backend/internal/infrastructure/persistence"
	"github.com/rsmnarts/codebase-golang-backend/internal/usecase"
	"github.com/rsmnarts/codebase-golang-backend/pkg/config"
	"github.com/rsmnarts/codebase-golang-backend/pkg/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	// Setup middleware
	middleware.Setup(app)

	// Initialize dependencies
	userRepo := persistence.NewInMemoryUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := http.NewUserHandler(userUseCase)

	// Setup routes
	http.SetupRoutes(app, userHandler)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	log.Printf("Server starting on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
