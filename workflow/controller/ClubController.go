package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/workflow/models"
)

// CreateClub godoc
// @Tags Club
// @Summary Create Club
// @Description Create New Club
// @ID createClub
// @Accept json
// @Produce json
// @Param requestBody body models.Tbl_club true "User credentials in JSON format"
// @Success 201 {object} models.Tbl_club
// @Security Bearer
// @Router /api/workflow/club [post]
func CreateClub(c *fiber.Ctx) error {
	Club := new(models.Tbl_club) // Pastikan ini adalah pointer ke struct yang benar

	if err := c.BodyParser(Club); err != nil {
		// Jika terjadi error, kirim response dengan status 400 dan pesan error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	// Lanjutkan dengan proses penyimpanan ke database
	db := database.GetDB()
	result := db.Create(Club)

	// Jika terjadi error saat menyimpan, kirim response dengan status 500
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create club: " + result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    Club,
	})
}

// GetAllClub godoc
// @Tags Club
// @Summary Get all Club
// @Description Get details of all club
// @ID get-all-Club
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Tbl_club
// @Security Bearer
// @Router /api/workflow/club [get]
func GetAllClubs(c *fiber.Ctx) error {
	var Club []models.Tbl_club

	db := database.GetDB()
	result := db.Find(&Club)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Club,
	})
}

// @Tags Club
// @Summary Get club by ID
// @Description Get a club by ID
// @ID get-club-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Club ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Club tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/club/{id} [get]
func GetClubById(c *fiber.Ctx) error {
	db := database.GetDB()

	ClubID := c.Params("id")

	var Club models.Tbl_club
	data := db.First(&Club, ClubID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Turnament tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Club,
	})
}

// @Tags Club
// @Summary Get club by ID
// @Description Get a club by ID
// @ID get-club-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Club ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Club tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/club/{id} [get]
func UpdateClubById(c *fiber.Ctx) error {
	db := database.GetDB()

	ClubID := c.Params("id")

	var Club models.Tbl_club

	if err := db.First(&Club, ClubID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Club tidak ditemukan",
		})
	}
	updatedClub := new(models.Tbl_club)
	if err := c.BodyParser(updatedClub); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	Club.Name = updatedClub.Name
	Club.Hometown = updatedClub.Hometown

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&Club).Updates(&Club)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan Club",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    Club,
	})
}

// @Tags Club
// @Summary Delete club by ID
// @Description Delete a club by ID
// @ID delete-club-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Club ID"
// @Success 200 {object} map[string]interface{} "Berhasil Menghapus Data"
// @Failure 404 {object} map[string]interface{} "Club tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menghapus data Club"
// @Security Bearer
// @Router /api/workflow/club/{id} [delete]
func DeleteClubById(c *fiber.Ctx) error {
	db := database.GetDB()

	ClubID := c.Params("id")

	var Club models.Tbl_club
	data := db.First(&Club, ClubID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Club tidak ditemukan",
		})
	}
	if err := db.Delete(&Club).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data Club",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data Club",
	})
}
