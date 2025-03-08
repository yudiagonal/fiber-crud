package repositories

import (
	"fiber-crud/database"
	"fiber-crud/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (models.User, error) {
	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)

	// Simpan pengguna ke database
	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

// func UpdateUser(id uint, updateUser models.User) (*models.User, error) {
// 	var user models.User
// 	if err := database.DB.First(&user, id).Error; err != nil {
// 		return nil, err
// 	}

// 	user.Username = updateUser.Username
// 	if updateUser.Password != ""{
// 		hash
// 	}
// }

func FindUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
