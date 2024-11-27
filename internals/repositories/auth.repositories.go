package repository

import (
	"gofiber-boiler/config"
	"gofiber-boiler/database"
	"gofiber-boiler/internals/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthRepository interface {
	GenerateToken(user *model.User, c *fiber.Ctx) (string, error)
	RenewToken(c *fiber.Ctx) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository() AuthRepository {
	return &authRepository{db: database.DB}
}

func (r *authRepository) GenerateToken(user *model.User, c *fiber.Ctx) (string, error) {
	day := time.Hour * 24

	// Create token
	clams := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	})
	// Generate encoded token and send it as response.
	token, err := clams.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return token, nil
}

func (r *authRepository) RenewToken(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	user := new(model.User)
	t, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return err
	}

	payload := t.Claims.(jwt.MapClaims)
	user.ID = uint(payload["ID"].(float64))
	user.Email = payload["email"].(string)
	user.Name = payload["name"].(string)

	if _, err := r.GenerateToken(user, c); err != nil {
		return err
	}

	return nil
}
