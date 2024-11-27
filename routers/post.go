package router

import (
	postHandler "gofiber-boiler/internals/handlers"
	"gofiber-boiler/internals/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(router fiber.Router) {
	post := router.Group("/posts")
	// Create a post
	post.Post("/", postHandler.CreatePost)
	// Read all posts
	post.Get("/", middlewares.AuthMiddleware, postHandler.GetPosts)
}
