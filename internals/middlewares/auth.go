package middlewares

import (
	repository "gofiber-boiler/internals/repositories"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Middleware JWT function
func AuthMiddleware(c *fiber.Ctx) error {

	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte("secret")},
		TokenLookup: "cookie:jwt",
		ContextKey:  "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return status 401 and failed authentication error.
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			if err := repository.NewAuthRepository().RenewToken(c); err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "InValid Token", "data": err})
			}
			return c.Next()
		},
	})(c)
}
