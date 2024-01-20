package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/master/models"
	"go-scoresheet/master/service"
	"strconv"
)

type RoleController struct {
	roleService service.RoleService
}

func NewRoleController(roleService service.RoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

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
func (c *RoleController) CreateRole(ctx *fiber.Ctx) error {
	role := new(models.Role)

	if err := ctx.BodyParser(role); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	if err := c.roleService.CreateRole(role); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create role: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    role,
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
func (c *RoleController) GetAllRoles(ctx *fiber.Ctx) error {
	roles, err := c.roleService.GetAllRoles()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch roles",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    roles,
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
func (c *RoleController) GetRoleById(ctx *fiber.Ctx) error {
	roleID := ctx.Params("id")
	id, err := strconv.ParseUint(roleID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Role ID",
		})
	}

	role, err := c.roleService.GetRoleById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Role not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    role,
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
func (c *RoleController) UpdateRoleById(ctx *fiber.Ctx) error {
	roleID := ctx.Params("id")
	id, err := strconv.ParseUint(roleID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Role ID",
		})
	}

	role, err := c.roleService.GetRoleById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Role not found",
		})
	}

	updatedRole := new(models.Role)
	if err := ctx.BodyParser(updatedRole); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	role.Code = updatedRole.Code
	role.Name = updatedRole.Name
	role.Guard = updatedRole.Guard
	role.Tag = updatedRole.Tag

	if err := c.roleService.UpdateRole(role); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Role",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated Role",
		"data":    role,
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
func (c *RoleController) DeleteRoleById(ctx *fiber.Ctx) error {
	roleID := ctx.Params("id")
	id, err := strconv.ParseUint(roleID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Role ID",
		})
	}

	if err := c.roleService.DeleteRole(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Role",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted Role",
	})
}
