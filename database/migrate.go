package database

import "fiber-crud/models"

func Migrate() {
	DB.AutoMigrate(&models.Product{}, &models.User{})
}
