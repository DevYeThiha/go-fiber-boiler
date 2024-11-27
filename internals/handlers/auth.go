package handler

import (
	"gofiber-boiler/database"
	"gofiber-boiler/internals/model"
	repository "gofiber-boiler/internals/repositories"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
)

type AuthHandler struct {
	AuthRepo repository.AuthRepository
}

func NewAuthHandler(authRepo repository.AuthRepository) *AuthHandler {
	return &AuthHandler{AuthRepo: authRepo}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	db := database.DB
	loginUser := new(model.User)
	var user model.User

	// Store the body in the post and return error if encountered
	if err := c.BodyParser(loginUser); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Find the user by credentials
	if err := db.Where("email = ?", loginUser.Email).Find(&user).Error; err != nil {
		log.Println("err fetching user")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	token, err := h.AuthRepo.GenerateToken(&user, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": struct {
		Token string `json:"token"`
	}{Token: token}})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	// Store the body in the post and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the note
	user.UUID = uuid.New().String()
	// post.CreatedAt = time.Now()
	// Create the Note and return error if encountered
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created User", "data": user})
}
