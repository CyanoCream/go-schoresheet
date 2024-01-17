package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/master/models"
)

// CreateUserRole godoc
// @Tags User Roles
// @Summary Create User Roles
// @Description Create New User Role
// @ID createUserRole
// @Accept json
// @Produce json
// @Param requestBody body models.UserRole true "User credentials in JSON format"
// @Success 201 {object} models.UserRole
// @Security Bearer
// @Router /api/user-role [post]
func CreateUserRole(c *fiber.Ctx) error {
	UserRole := new(models.UserRole) // Pastikan ini adalah pointer ke struct yang benar

	// Parse body ke struct UserRole
	if err := c.BodyParser(UserRole); err != nil {
		// Jika terjadi error, kirim response dengan status 400 dan pesan error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	// Lanjutkan dengan proses penyimpanan ke database
	db := database.GetDB()
	result := db.Create(UserRole)

	// Jika terjadi error saat menyimpan, kirim response dengan status 500
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user role: " + result.Error.Error(),
		})
	}

	// Jika berhasil, kirim response dengan status 201 dan data Role yang telah dibuat
	return c.Status(fiber.StatusCreated).JSON(c.JSON(fiber.Map{
		"message": "success",
		"data":    UserRole,
	}))
}

// GetAllUserRoles godoc
// @Tags User Roles
// @Summary Get all User Roles
// @Description Get details of all user roles
// @ID get-all-UserRoles
// @Accept  json
// @Produce  json
// @Success 200 {array} models.UserRole
// @Security Bearer
// @Router /api/user-role [get]
func GetAllUserRoles(c *fiber.Ctx) error {
	var UserRole []models.UserRole

	db := database.GetDB()
	result := db.Find(&UserRole)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    UserRole,
	})
}

// @Tags User Roles
// @Summary Get user role by ID
// @Description Get a user role by ID
// @ID get-user-role-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "User Role ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "User Role tidak ditemukan"
// @Security Bearer
// @Router /api/user-role/{id} [get]
func GetUserRoleById(c *fiber.Ctx) error {
	db := database.GetDB()

	UserRoleID := c.Params("id")

	var UserRole models.UserRole
	data := db.First(&UserRole, UserRoleID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Role tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    UserRole,
	})
}

// @Tags User Roles
// @Summary Get user role by ID
// @Description Get a user role by ID
// @ID get-user-role-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "User Role ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "User Role tidak ditemukan"
// @Security Bearer
// @Router /api/user-role/{id} [get]
func UpdateUserRoleById(c *fiber.Ctx) error {
	db := database.GetDB()

	UserRoleID := c.Params("id")

	var UserRole models.UserRole

	if err := db.First(&UserRole, UserRoleID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Role tidak ditemukan",
		})
	}
	updatedUserRole := new(models.UserRole)
	if err := c.BodyParser(updatedUserRole); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	UserRole.RoleCode = updatedUserRole.RoleCode
	UserRole.UserID = updatedUserRole.UserID

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&UserRole).Updates(&UserRole)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan Permission",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    UserRole,
	})
}

// @Tags User Roles
// @Summary Delete user role by ID
// @Description Delete a user role by ID
// @ID delete-user-role-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "User Role ID"
// @Success 200 {object} map[string]interface{} "Berhasil Menghapus Data"
// @Failure 404 {object} map[string]interface{} "User Role tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menghapus data user role"
// @Security Bearer
// @Router /api/userroles/{id} [delete]
func DeleteUserRoleById(c *fiber.Ctx) error {
	db := database.GetDB()

	UserRoleID := c.Params("id")

	var UserRole models.UserRole
	data := db.First(&UserRole, UserRoleID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data User Role tidak ditemukan",
		})
	}
	if err := db.Delete(&UserRole).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data User Role",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data User Role",
	})
}

func GetUserRolesByID(userID uint) []models.UserRole {
	var userRoles []models.UserRole

	// GORM database connection
	db := database.GetDB() // Mengambil instance GORM dari database yang telah Anda inisialisasi

	// Query user roles based on user ID
	db.Where("user_id = ?", userID).Find(&userRoles)

	return userRoles
}
