package interfaces

import (
	"fiber-crud/models"
)

type ProductRepository interface {
	GetAllProducts(page, limit int) ([]models.Product, int64, error)
	GetProductByID(id int) *models.Product
	CreateProduct(product models.Product) models.Product
	UpdateProduct(id int, updatedProduct models.Product) *models.Product
	DeleteProduct(id int) bool
}
