package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func (app *App) jwtMiddleware(c *fiber.Ctx) error {
	stringToken := c.Get("Authorization")
	if stringToken == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.config.Secret), nil
	})
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	c.Locals("user", token)
	return c.Next()
}
