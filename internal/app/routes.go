package app

import (
	"github.com/gofiber/fiber/v2"
)

func (app *App) Routes(fApp *fiber.App) {
	fApp.Post("/AddUser", app.jwtMiddleware, app.AddUser)
	fApp.Post("/Login", app.Login)
	fApp.Get("/User/:id", app.jwtMiddleware, app.GetUserById)
	fApp.Get("/User", app.jwtMiddleware, app.GetAllUsers)
	fApp.Get("/User/Delete/:id", app.jwtMiddleware, app.DeleteUser)
	fApp.Post("/User/Update", app.jwtMiddleware, app.UpdateUser)
}
