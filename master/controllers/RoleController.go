package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/master/models"
)

// CreateRole godoc
// @Tags Roles
// @Summary Create Roles
// @Description Create New Role
// @ID createRole
// @Accept json
// @Produce json
// @Param requestBody body models.Role true "User credentials in JSON format"
// @Success 201 {object} models.Role
// @Security Bearer
// @Router /api/role [post]
func CreateRole(c *fiber.Ctx) error {
	Role := new(models.Role) // Pastikan ini adalah pointer ke struct yang benar

	// Parse body ke struct Role
	if err := c.BodyParser(Role); err != nil {
		// Jika terjadi error, kirim response dengan status 400 dan pesan error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	// Lanjutkan dengan proses penyimpanan ke database
	db := database.GetDB()
	result := db.Create(Role)

	// Jika terjadi error saat menyimpan, kirim response dengan status 500
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create role: " + result.Error.Error(),
		})
	}

	// Jika berhasil, kirim response dengan status 201 dan data Role yang telah dibuat
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    Role,
	})
}

// GetAllRoles godoc
// @Tags Roles
// @Summary Get all Roles
// @Description Get details of all roles
// @ID get-all-Roles
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Role
// @Security Bearer
// @Router /api/role [get]
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

// @Tags Roles
// @Summary Get role by ID
// @Description Get a role by ID
// @ID get-role-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Role tidak ditemukan"
// @Security Bearer
// @Router /api/role/{id} [get]
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

// @Tags Roles
// @Summary Get role by ID
// @Description Get a role by ID
// @ID get-role-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Role tidak ditemukan"
// @Security Bearer
// @Router /api/role/{id} [get]
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

// @Tags Roles
// @Summary Delete role by ID
// @Description Delete a role by ID
// @ID delete-role-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Role ID"
// @Success 200 {object} map[string]interface{} "Berhasil Menghapus Data"
// @Failure 404 {object} map[string]interface{} "Role tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menghapus data role"
// @Security Bearer
// @Router /api/roles/{id} [delete]
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
