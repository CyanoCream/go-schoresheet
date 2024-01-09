package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/master/models"
)

func CreateRole(c *fiber.Ctx) error {
	Role := new(models.Role)

	if err := c.BodyParser(Role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	db := database.GetDB()
	result := db.Create(Role)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create Role",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(Role)
}

// Get all users
func GetAllRoles(c *fiber.Ctx) error {
	var Role []models.Role

	db := database.GetDB()
	result := db.Find(&Role)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Role,
	})
}

func GetRoleById(c *fiber.Ctx) error {
	db := database.GetDB()

	RoleID := c.Params("id")

	var Role models.Role
	data := db.First(&Role, RoleID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Role tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Role,
	})
}

func UpdateRoleById(c *fiber.Ctx) error {
	db := database.GetDB()

	RoleID := c.Params("id")

	var Role models.Role

	if err := db.First(&Role, RoleID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Role tidak ditemukan",
		})
	}
	updatedRole := new(models.Role)
	if err := c.BodyParser(updatedRole); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	Role.Code = updatedRole.Code
	Role.Name = updatedRole.Name
	Role.Guard = updatedRole.Guard
	Role.Tag = updatedRole.Tag

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&Role).Updates(&Role)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan Permission",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    Role,
	})
}
func DeleteRoleById(c *fiber.Ctx) error {
	db := database.GetDB()

	RoleID := c.Params("id")

	var Role models.Role
	data := db.First(&Role, RoleID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Role tidak ditemukan",
		})
	}
	if err := db.Delete(&Role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data Role",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data Role",
	})
}
