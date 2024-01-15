package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/master/models"
)

// CreatePermissionRole godoc
// @Tags Permission Roles
// @Summary Create Permission Role
// @Description Create New Permission Role
// @ID CreatePermissionRole
// @Accept json
// @Produce json
// @Param requestBody body models.PermissionRole true "Permissioan Role credentials in JSON format"
// @Success 201 {object} models.PermissionRole
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission-role [post]
func CreatePermissionRole(c *fiber.Ctx) error {
	PermissionRole := new(models.PermissionRole)

	if err := c.BodyParser(PermissionRole); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	db := database.GetDB()
	result := db.Create(PermissionRole)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create Permission Role",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(PermissionRole)
}

// GetAllPermissionRoles godoc
// @Tags Permission Roles
// @Summary Get All Permission Role
// @Description Get All Permission Role
// @ID GetAllPermissionRole
// @Accept json
// @Produce json
// @Success 201 {object} models.PermissionRole
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission-role [get]
func GetAllPermissionRoles(c *fiber.Ctx) error {
	var PermissionRole []models.PermissionRole

	db := database.GetDB()
	result := db.Find(&PermissionRole)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions Role",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    PermissionRole,
	})
}

// GetPermissionRoleById godoc
// @Tags Permission Roles
// @Summary Get Permission Role by ID
// @Description Get Permission Role by ID
// @ID GetPermissionRoleById
// @Accept json
// @Produce json
// @Param id path string true "PermissionRole ID"
// @Success 201 {object} models.PermissionRole
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission-role/{id} [get]
func GetPermissionRoleById(c *fiber.Ctx) error {
	db := database.GetDB()

	PermissionRoleID := c.Params("id")

	var PermissionRole models.PermissionRole
	data := db.First(&PermissionRole, PermissionRoleID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Permission Role tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    PermissionRole,
	})
}

// UpdatePermissionRoleById godoc
// @Tags Permission Roles
// @Summary Update Permission Role by ID
// @Description Update Permission Role by ID
// @ID UpdatePermissionRoleById
// @Accept json
// @Produce json
// @Param id path string true "PermissionRole ID"
// @Success 201 {object} models.PermissionRole
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission-role/{id} [post]
func UpdatePermissionRoleById(c *fiber.Ctx) error {
	db := database.GetDB()

	PermissionRoleID := c.Params("id")

	var PermissionRole models.PermissionRole

	if err := db.First(&PermissionRole, PermissionRoleID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Permission Role tidak ditemukan",
		})
	}
	updatedPermissionRole := new(models.PermissionRole)
	if err := c.BodyParser(updatedPermissionRole); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	PermissionRole.RoleCode = updatedPermissionRole.RoleCode
	PermissionRole.PermissionCode = updatedPermissionRole.PermissionCode

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&PermissionRole).Updates(&PermissionRole)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan Permission",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    PermissionRole,
	})
}

// DeletePermissionRoleById godoc
// @Tags Permission Roles
// @Summary Delete Permission Role by ID
// @Description Delete Permission Role by ID
// @ID DeletePermissionRoleById
// @Accept json
// @Produce json
// @Param id path string true "PermissionRole ID"
// @Success 201 {object} models.PermissionRole
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission-role/{id} [delete]
func DeletePermissionRoleById(c *fiber.Ctx) error {
	db := database.GetDB()

	PermissionRoleID := c.Params("id")

	var PermissionRole models.PermissionRole
	data := db.First(&PermissionRole, PermissionRoleID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Permission Role tidak ditemukan",
		})
	}
	if err := db.Delete(&PermissionRole).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data Permission Role",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data permission Role",
	})
}
