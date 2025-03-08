package interfaces

import (
	"fiber-crud/models"
)

type ProductRepository interface {
	GetAllProduct() []models.Product
	GetProductByID(id int) *models.Product
	CreateProduct(product models.Product) models.Product
	UpdateProduct(id int, updatedProduct models.Product) *models.Product
	DeleteProduct(id int) bool
}
