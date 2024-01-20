package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/master/models"
	"go-scoresheet/master/service"
	"strconv"
)

type UserRoleController struct {
	userRoleService service.UserRoleService
}

func NewUserRoleController(userRoleService service.UserRoleService) *UserRoleController {
	return &UserRoleController{
		userRoleService: userRoleService,
	}
}

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
func (c *UserRoleController) CreateUserRole(ctx *fiber.Ctx) error {
	userRole := new(models.UserRole)

	if err := ctx.BodyParser(userRole); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	if err := c.userRoleService.CreateUserRole(userRole); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user role: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    userRole,
	})

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
func (c *UserRoleController) GetAllUserRoles(ctx *fiber.Ctx) error {
	userRoles, err := c.userRoleService.GetAllUserRoles()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch user roles",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    userRoles,
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
func (c *UserRoleController) GetUserRoleById(ctx *fiber.Ctx) error {
	userRoleID := ctx.Params("id")
	id, err := strconv.ParseUint(userRoleID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid User Role ID",
		})
	}

	userRole, err := c.userRoleService.GetUserRoleById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Role not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    userRole,
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
func (c *UserRoleController) UpdateUserRoleById(ctx *fiber.Ctx) error {
	userRoleID := ctx.Params("id")
	id, err := strconv.ParseUint(userRoleID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid User Role ID",
		})
	}

	userRole, err := c.userRoleService.GetUserRoleById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Role not found",
		})
	}

	updatedUserRole := new(models.UserRole)
	if err := ctx.BodyParser(updatedUserRole); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	userRole.RoleCode = updatedUserRole.RoleCode
	userRole.UserID = updatedUserRole.UserID

	if err := c.userRoleService.UpdateUserRole(userRole); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update User Role",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated User Role",
		"data":    userRole,
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
func (c *UserRoleController) DeleteUserRoleById(ctx *fiber.Ctx) error {
	userRoleID := ctx.Params("id")
	id, err := strconv.ParseUint(userRoleID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid User Role ID",
		})
	}

	if err := c.userRoleService.DeleteUserRole(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete User Role",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted User Role",
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
