package router

import (
	handler "gofiber-boiler/internals/handlers"
	repository "gofiber-boiler/internals/repositories"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	post := router.Group("/auth")
	authRepo := repository.NewAuthRepository()
	authHandler := handler.NewAuthHandler(authRepo)
	// Create a post
	post.Post("/login", authHandler.Login)
	// Read all posts
	post.Post("/register", authHandler.Register)
}
