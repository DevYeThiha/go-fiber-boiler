package handler

import (
	"gofiber-boiler/database"
	"gofiber-boiler/internals/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetPost func gets all existing notes
// @Description Get all existing notes
// @Tags Notes
// @Accept json
// @Produce json
// @Success 200 {array} model.Note
// @router /api/note [get]
func GetPosts(c *fiber.Ctx) error {
	db := database.DB
	var post []model.Post

	// find all post in the database
	db.Find(&post)

	// If no note is present return an error
	if len(post) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Post present", "data": nil})
	}

	// Else return notes
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": post})
}

// CreatePost func create a note
// @Description Create a Note
// @Tags Notes
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param sub_title body string true "SubTitle"
// @Param text body string true "Text"
// @Success 200 {object} model.Note
// @router /api/note [post]
func CreatePost(c *fiber.Ctx) error {
	db := database.DB
	post := new(model.Post)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(post)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the note
	post.UUID = uuid.New().String()
	// post.CreatedAt = time.Now()
	// Create the Note and return error if encountered
	err = db.Create(&post).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": post})
}
