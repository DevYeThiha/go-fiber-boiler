package main

import (
	"gofiber-boiler/config"
	"gofiber-boiler/database"
	router "gofiber-boiler/routers"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOriginsFunc: func(origin string) bool {
			return strings.Contains(origin, ":://localhost")
		},
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	port := config.Config("PORT")

	// Setup the router
	router.SetupRoutes(app)
	app.Listen(":" + port)
}
