package app

import (
	"context"
	"time"

	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/entities"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/models"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/repository"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (app *App) Login(c *fiber.Ctx) error {
	loginModel := &models.LoginModel{}
	ctx := context.Background()
	if err := c.BodyParser(&loginModel); err != nil {
		return c.Status(fiber.StatusOK).JSON(utils.StatusFail("Error occurred while generating token"))
	}
	repo := repository.NewRepository[entities.User](app.db)
	user := repo.Get(&entities.User{Username: loginModel.Username, Password: loginModel.Password}, ctx)
	if user.ID > 0 {
		claims := jwt.MapClaims{
			"id":      user.ID,
			"name":    user.Name,
			"surname": user.Surname,
			"email":   user.Email,
			"exp":     time.Now().Add(time.Minute * time.Duration(app.config.Expire)).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(app.config.Secret))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while generating token"))
		}
		mapToken := map[string]string{
			"token": tokenString,
		}
		return c.Status(fiber.StatusOK).JSON(utils.StatusOK(mapToken))
	}
	return c.Status(fiber.StatusUnauthorized).JSON(utils.StatusUnauthorized("Invalid username or password"))
}
