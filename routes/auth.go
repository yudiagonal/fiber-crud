package routes

import (
	"fiber-crud/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/register", controllers.CreateUser)
	app.Post("/login", controllers.Login)
}
