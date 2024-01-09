package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/master/models"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	db := database.GetDB()
	result := db.Create(user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Get all users
func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	db := database.GetDB()
	result := db.Find(&users)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}

func GetUserById(c *fiber.Ctx) error {
	db := database.GetDB()

	userID := c.Params("id")

	var user models.User
	data := db.First(&user, userID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UpdateUserById(c *fiber.Ctx) error {
	db := database.GetDB()

	userID := c.Params("id")

	var user models.User

	if err := db.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	}
	updatedUser := new(models.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	user.Fullname = updatedUser.Fullname
	user.Email = updatedUser.Email
	user.Username = updatedUser.Username
	user.Password = updatedUser.Password

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&user).Updates(&user)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    user,
	})
}
func DeleteUserById(c *fiber.Ctx) error {
	db := database.GetDB()

	userID := c.Params("id")

	var user models.User
	data := db.First(&user, userID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	}
	if err := db.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data pengguna",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data",
	})
}
