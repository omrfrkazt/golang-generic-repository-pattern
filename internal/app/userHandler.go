package app

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/entities"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/models"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/repository"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/utils"
)

func (app *App) AddUser(c *fiber.Ctx) error {
	ctx := context.Background()
	user := &models.UserModel{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Invalid request"))
	}
	repo := repository.NewRepository[entities.User](app.db)
	err := repo.Add(utils.UserToEntity(user), ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while adding user"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK("User added successfully"))
}

func (app *App) UpdateUser(c *fiber.Ctx) error {
	ctx := context.Background()
	user := &models.UserModel{}
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Invalid request"))
	}
	repo := repository.NewRepository[entities.User](app.db)
	err := repo.Update(utils.UserToEntity(user), ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while updating user"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK("User updated successfully"))
}

func (app *App) GetUserById(c *fiber.Ctx) error {
	ctx := context.Background()
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Invalid request"))
	}
	repo := repository.NewRepository[entities.User](app.db)
	user, err := repo.GetById(id, ctx)
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(utils.StatusNotFound("User not found"))
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while getting user"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK(utils.UserToModel(user)))
}

func (app *App) GetAllUsers(c *fiber.Ctx) error {
	ctx := context.Background()
	repo := repository.NewRepository[entities.User](app.db)
	users, err := repo.Where(&entities.User{IsActive: true}, ctx)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while getting users"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK(utils.UsersToModel(users)))
}

func (app *App) DeleteUser(c *fiber.Ctx) error {
	ctx := context.Background()
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Invalid request"))
	}
	repo := repository.NewRepository[entities.User](app.db)
	err = repo.Delete(id, ctx)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.StatusFail("Error occurred while deleting user"))
	}
	return c.Status(fiber.StatusOK).JSON(utils.StatusOK("User deleted successfully"))
}
