package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/master/helpers"
	"go-scoresheet/master/models"
	"go-scoresheet/middleware"
	"net/http"
)

// @Summary Login user
// @Description Logs in a user and returns an authentication token
// @ID loginUser
// @Tags Auth
// @Accept json
// @Produce json
// @Success 201 {object} middleware.JWT
// @Param requestBody body middleware.LoginField true "User credentials in JSON format"
// @Router /api/login [post]
func LoginUser(c *fiber.Ctx) error {
	db := database.GetDB()
	user := models.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": err.Error(),
		})
	}

	password := user.Password

	err = db.Where("username = ?", user.Username).Take(&user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": "invalid username or password",
		})
	}

	if !helpers.PasswordValid(user.Password, password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusText(http.StatusBadRequest),
			"message": "invalid username or password",
		})
	}

	// Check if the user is already active
	var session middleware.Session
	if db != nil {
		db.First(&session, "user_id = ?", user.ID)
	}

	if session.UserID != 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "User sedang aktif",
		})
	}
	userRoles := GetUserRolesByID(user.ID)

	var roleCodes []string
	for _, userRole := range userRoles {
		roleCodes = append(roleCodes, userRole.RoleCode)
	}

	token, err := helpers.GenerateToken(user.ID, user.Email, roleCodes)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusText(http.StatusInternalServerError),
			"message": err.Error(),
		})
	}

	// Save the session after the token is successfully generated
	err = saveSession(c, int(user.ID), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusText(http.StatusInternalServerError),
			"message": "Failed to create session",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

func saveSession(c *fiber.Ctx, userId int, token string) error {
	db := database.GetDB()
	sessionData := middleware.Session{
		UserID: userId,
		Token:  token,
	}
	if db != nil {
		result := db.Create(&sessionData)

		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create session",
			})
		}
	}
	return nil
}

// @Summary Logout user
// @Description Session Logout
// @ID logoutUser
// @Tags Auth
// @Accept json
// @Produce json
// @Router /api/logout [DELETE]
func DeleteSessionByToken(c *fiber.Ctx) error {
	// Membuka koneksi ke database
	db := database.GetDB()

	// Ambil token dari request JSON
	var req struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	// Mencari sesi berdasarkan token
	var session middleware.Session
	if err := db.Unscoped().Where("token = ?", req.Token).First(&session).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Session not found",
		})
	}

	// Menghapus sesi dari database
	if err := db.Unscoped().Delete(&session).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete session",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Session deleted successfully",
	})
}
