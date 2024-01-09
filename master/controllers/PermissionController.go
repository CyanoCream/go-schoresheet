package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/master/models"
)

func CreatePermission(c *fiber.Ctx) error {
	permission := new(models.Permission)

	if err := c.BodyParser(permission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	db := database.GetDB()
	result := db.Create(permission)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create permission",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(permission)
}

// Get all users
func GetAllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	db := database.GetDB()
	result := db.Find(&permissions)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    permissions,
	})
}

func GetPermissionById(c *fiber.Ctx) error {
	db := database.GetDB()

	permissionID := c.Params("id")

	var permission models.Permission
	data := db.First(&permission, permissionID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Permission tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    permission,
	})
}

func UpdatePermissionById(c *fiber.Ctx) error {
	db := database.GetDB()

	permissionID := c.Params("id")

	var permission models.Permission

	if err := db.First(&permission, permissionID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Permission tidak ditemukan",
		})
	}
	updatedPermission := new(models.Permission)
	if err := c.BodyParser(updatedPermission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	permission.Code = updatedPermission.Code
	permission.Name = updatedPermission.Name
	permission.GuardName = updatedPermission.GuardName
	permission.Description = updatedPermission.Description
	permission.Action = updatedPermission.Action
	permission.Module = updatedPermission.Module
	permission.Tag = updatedPermission.Tag
	permission.IsActive = updatedPermission.IsActive

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&permission).Updates(&permission)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan Permission",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    permission,
	})
}
func DeletePermissionById(c *fiber.Ctx) error {
	db := database.GetDB()

	permissionID := c.Params("id")

	var permission models.Permission
	data := db.First(&permission, permissionID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Permission tidak ditemukan",
		})
	}
	if err := db.Delete(&permission).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data permission",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data Permission",
	})
}
