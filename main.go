package main

import (
	"fiber-crud/database"
	"fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database
	database.Connect()
	database.Migrate()

	// Initialize Fiber app
	app := fiber.New()

	// // Serve static files (HTML, CSS, JS)
	// app.Use("/", filesystem.New(filesystem.Config{
	// 	Root: http.Dir("./views"),
	// }))

	// Setup routes

	routes.SetupAuthRoutes(app)
	routes.SetupRoutes(app)

	// Start server
	app.Listen(":3000")
}
