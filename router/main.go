package router

import (
	"github.com/Sittikorn-off/Challenge_go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	challenge := app.Group("/challenge")
	challenge.Get("/c1", handlers.Challenge1)
	challenge.Post("/c2", handlers.Challenge2)
	challenge.Get("/c3", handlers.Challenge3)

}
