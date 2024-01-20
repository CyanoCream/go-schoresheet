package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/master/models"
	"go-scoresheet/master/service"
	"strconv"
)

type PermissionRoleController struct {
	permissionRoleService service.PermissionRoleService
}

func NewPermissionRoleController(permissionRoleService service.PermissionRoleService) *PermissionRoleController {
	return &PermissionRoleController{
		permissionRoleService: permissionRoleService,
	}
}

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
func (c *PermissionRoleController) CreatePermissionRole(ctx *fiber.Ctx) error {
	permissionRole := new(models.PermissionRole)

	if err := ctx.BodyParser(permissionRole); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	if err := c.permissionRoleService.CreatePermissionRole(permissionRole); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create Permission Role",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(permissionRole)
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
func (c *PermissionRoleController) GetAllPermissionRoles(ctx *fiber.Ctx) error {
	permissionRoles, err := c.permissionRoleService.GetAllPermissionRoles()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permission roles",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    permissionRoles,
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
func (c *PermissionRoleController) GetPermissionRoleById(ctx *fiber.Ctx) error {
	permissionRoleID := ctx.Params("id")
	id, err := strconv.ParseUint(permissionRoleID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Permission Role ID",
		})
	}

	permissionRole, err := c.permissionRoleService.GetPermissionRoleById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Permission Role not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    permissionRole,
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
func (c *PermissionRoleController) UpdatePermissionRoleById(ctx *fiber.Ctx) error {
	permissionRoleID := ctx.Params("id")
	id, err := strconv.ParseUint(permissionRoleID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Permission Role ID",
		})
	}

	permissionRole, err := c.permissionRoleService.GetPermissionRoleById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Permission Role not found",
		})
	}

	updatedPermissionRole := new(models.PermissionRole)
	if err := ctx.BodyParser(updatedPermissionRole); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	permissionRole.RoleCode = updatedPermissionRole.RoleCode
	permissionRole.PermissionCode = updatedPermissionRole.PermissionCode

	if err := c.permissionRoleService.UpdatePermissionRole(permissionRole); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Permission Role",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated Permission Role",
		"data":    permissionRole,
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
func (c *PermissionRoleController) DeletePermissionRoleById(ctx *fiber.Ctx) error {
	permissionRoleID := ctx.Params("id")
	id, err := strconv.ParseUint(permissionRoleID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Permission Role ID",
		})
	}

	if err := c.permissionRoleService.DeletePermissionRole(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Permission Role",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted Permission Role",
	})
}
