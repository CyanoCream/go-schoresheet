package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/master/models"
	"go-scoresheet/master/service"
	"strconv"
)

type PermissionController struct {
	permissionService service.PermissionService
}

func NewPermissionController(permissionService service.PermissionService) *PermissionController {
	return &PermissionController{
		permissionService: permissionService,
	}
}

// CreatePermission godoc
// @Tags Permissions
// @Summary Create Permission
// @Description Create New Permission
// @ID CreatePermission
// @Accept json
// @Produce json
// @Param requestBody body models.Permission true "Permissioan credentials in JSON format"
// @Success 201 {object} models.Permission
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission [post]
func (c *PermissionController) CreatePermission(ctx *fiber.Ctx) error {
	permission := new(models.Permission)

	if err := ctx.BodyParser(permission); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	if err := c.permissionService.CreatePermission(permission); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create permission",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(permission)
}

// GetAllPermission godoc
// @Tags Permissions
// @Summary Get All Permission
// @Description Get All Permission
// @ID GetAllPermission
// @Accept json
// @Produce json
// @Success 201 {object} models.Permission
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission [get]
func (c *PermissionController) GetAllPermissions(ctx *fiber.Ctx) error {
	permissions, err := c.permissionService.GetAllPermissions()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    permissions,
	})
}

// GetPermissionById godoc
// @Tags Permissions
// @Summary Get Permission by ID
// @Description Get Permission by ID
// @ID GetPermissionById
// @Accept json
// @Produce json
// @Param id path string true "Permission ID"
// @Success 201 {object} models.Permission
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission/{id} [get]
func (c *PermissionController) GetPermissionById(ctx *fiber.Ctx) error {
	permissionID := ctx.Params("id")
	id, err := strconv.ParseUint(permissionID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Permission ID",
		})
	}

	permission, err := c.permissionService.GetPermissionById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Permission not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    permission,
	})
}

// UpdatePermissionById godoc
// @Tags Permissions
// @Summary Update Permission by ID
// @Description Update Permission by ID
// @ID UpdatePermissionById
// @Accept json
// @Produce json
// @Param id path string true "Permission ID"
// @Success 201 {object} models.Permission
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission/{id} [post]
func (c *PermissionController) UpdatePermissionById(ctx *fiber.Ctx) error {
	permissionID := ctx.Params("id")
	id, err := strconv.ParseUint(permissionID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Permission ID",
		})
	}

	permission, err := c.permissionService.GetPermissionById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Permission not found",
		})
	}

	updatedPermission := new(models.Permission)
	if err := ctx.BodyParser(updatedPermission); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

	if err := c.permissionService.UpdatePermission(permission); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update permission",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated permission",
		"data":    permission,
	})
}

// DeletePermissionById godoc
// @Tags Permissions
// @Summary Delete Permission by ID
// @Description Delete Permission by ID
// @ID DeletePermissionById
// @Accept json
// @Produce json
// @Param id path string true "Permission ID"
// @Success 201 {object} models.Permission
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/permission/{id} [delete]
func (c *PermissionController) DeletePermissionById(ctx *fiber.Ctx) error {
	permissionID := ctx.Params("id")
	id, err := strconv.ParseUint(permissionID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Permission ID",
		})
	}

	if err := c.permissionService.DeletePermission(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete permission",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted permission",
	})
}
