package repositories

import (
	"fiber-crud/database"
	"fiber-crud/interfaces"
	"fiber-crud/models"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() interfaces.ProductRepository {
	return &ProductRepositoryImpl{}
}

// CreateProduct implements interfaces.ProductRepository.
func (p *ProductRepositoryImpl) CreateProduct(product models.Product) models.Product {
	database.DB.Create(&product)
	return product
}

// DeleteProduct implements interfaces.ProductRepository.
func (p *ProductRepositoryImpl) DeleteProduct(id int) bool {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return false
	}
	database.DB.Delete(&product)
	return true

}

// GetAllProduct implements interfaces.ProductRepository.
// GetAllProducts mengambil daftar produk dengan pagination
func (r *ProductRepositoryImpl) GetAllProducts(page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	// Hitung total data produk
	if err := database.DB.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Hitung offset berdasarkan halaman dan limit
	offset := (page - 1) * limit

	// Ambil data produk dengan pagination
	if err := database.DB.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// GetProductByID implements interfaces.ProductRepository.
func (p *ProductRepositoryImpl) GetProductByID(id int) *models.Product {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return nil
	}
	return &product
}

// UpdateProduct implements interfaces.ProductRepository.
func (p *ProductRepositoryImpl) UpdateProduct(id int, updatedProduct models.Product) *models.Product {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return nil
	}
	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	database.DB.Save(&product)
	return &product

}
