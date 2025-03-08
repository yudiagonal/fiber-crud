// routes/product_routes.go
package routes

import (
	"fiber-crud/controllers"
	"fiber-crud/interfaces"
	"fiber-crud/middlewares"
	"fiber-crud/repositories"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Initialize repository and controller
	var repo interfaces.ProductRepository = repositories.NewProductRepository()
	productController := controllers.NewProductController(repo)

	productsRoute := app.Group("/products", middlewares.AuthMiddleware)
	// Define routes
	productsRoute.Get("/", productController.GetAllProducts)
	productsRoute.Get("/:id", productController.GetProductByID)
	productsRoute.Post("/", productController.CreateProduct)
	productsRoute.Put("/:id", productController.UpdateProduct)
	productsRoute.Delete("/:id", productController.DeleteProduct)
}
