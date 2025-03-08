package middlewares

import (
	"fiber-crud/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Ambil token dari header Authorization
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized, token required",
		})
	}

	// Hapus prefix "Bearer " jika ada
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Verifikasi token
	token, err := utils.VerifyToken(tokenString)
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Ambil claims dari token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	// Ambil user_id dari claims
	userID, ok := claims["user_id"].(float64) // JWT menyimpan angka sebagai float64
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid user ID in token",
		})
	}

	// Simpan user ID ke context agar bisa digunakan di handler
	c.Locals("userID", uint(userID)) // Convert float64 ke uint

	// Lanjutkan ke handler berikutnya
	return c.Next()
}
