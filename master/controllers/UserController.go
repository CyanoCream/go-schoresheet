package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/master/models"
	"go-scoresheet/master/service"
	"strconv"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser godoc
// @Tags Users
// @Summary Create User
// @Description Create New User
// @ID createUser
// @Accept json
// @Produce json
// @Param requestBody body models.User true "User credentials in JSON format"
// @Success 201 {object} models.User
// @Router /api/users [post]
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	err := uc.userService.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetAllUsers godoc
// @Tags Users
// @Summary Get all users
// @Description Get details of all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /api/users [get]
// @Security ApiKeyAuth
// @Security Bearer
func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}

// @Tags Users
// @Summary Get user by ID
// @Description Get a user by ID
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "User tidak ditemukan"
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/users/{id} [get]
func (uc *UserController) GetUserById(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// @Tags Users
// @Summary Update user by ID
// @Description Update a user by ID
// @ID update-user-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} map[string]interface{} "Berhasil melakukan pembaruan"
// @Failure 404 {object} map[string]interface{} "User tidak ditemukan"
// @Failure 400 {object} map[string]interface{} "Invalid JSON"
// @Failure 500 {object} map[string]interface{} "Gagal melakukan pembaruan"
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/users/update/{id} [post]
func (c *UserController) UpdateUserById(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid User User ID",
		})
	}

	user, err := c.userService.GetUserById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Role not found",
		})
	}

	updatedUser := new(models.User)
	if err := ctx.BodyParser(updatedUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	user.Fullname = updatedUser.Fullname
	user.Email = updatedUser.Email
	user.Username = updatedUser.Username
	user.Password = updatedUser.Password

	err = c.userService.UpdateUser(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated user",
		"data":    user,
	})
}

// @Tags Users
// @Summary Delete user by ID
// @Description Delete a user by ID
// @ID delete-user-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{} "Berhasil Menghapus Data"
// @Failure 404 {object} map[string]interface{} "User tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menghapus data pengguna"
// @Security ApiKeyAuth
// @Security Bearer
// @Router /api/users/delete/{id} [delete]
func (c *UserController) DeleteUserById(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid User ID",
		})
	}

	if err := c.userService.DeleteUser(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete User",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted User",
	})
}
