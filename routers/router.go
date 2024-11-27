package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// app.Get("/swagger/*", swagger.Handler)
	// Group api calls with param '/api'
	// api := app.Group("/api", logger.New()
	api := app.Group("/api")

	// Setup note routes, can use same syntax to add routes for more models
	SetupPostRoutes(api)
	SetupAuthRoutes(api)
}
