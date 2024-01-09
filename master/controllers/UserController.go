package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/master/helpers"
	"go-scoresheet/master/models"
	"go-scoresheet/middleware"
	"net/http"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	db := database.GetDB()
	result := db.Create(user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Get all users
func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	db := database.GetDB()
	result := db.Find(&users)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch users",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}

func GetUserById(c *fiber.Ctx) error {
	db := database.GetDB()

	userID := c.Params("id")

	var user models.User
	data := db.First(&user, userID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UpdateUserById(c *fiber.Ctx) error {
	db := database.GetDB()

	userID := c.Params("id")

	var user models.User

	if err := db.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	}
	updatedUser := new(models.User)
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	user.Fullname = updatedUser.Fullname
	user.Email = updatedUser.Email
	user.Username = updatedUser.Username
	user.Password = updatedUser.Password

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&user).Updates(&user)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    user,
	})
}
func DeleteUserById(c *fiber.Ctx) error {
	db := database.GetDB()

	userID := c.Params("id")

	var user models.User
	data := db.First(&user, userID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	}
	if err := db.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data pengguna",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data",
	})
}

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

	token, err := helpers.GenerateToken(user.ID, user.Email)
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
