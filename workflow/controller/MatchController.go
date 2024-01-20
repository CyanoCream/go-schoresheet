package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/workflow/models"
)

// CreateMatch godoc
// @Tags Match
// @Summary Create Match
// @Description Create New Match
// @ID createMatch
// @Accept json
// @Produce json
// @Param requestBody body models.Tbl_match true "User credentials in JSON format"
// @Success 201 {object} models.Tbl_match
// @Security Bearer
// @Router /api/workflow/match [post]
func CreateMatch(c *fiber.Ctx) error {
	Match := new(models.Tbl_match) // Pastikan ini adalah pointer ke struct yang benar

	if err := c.BodyParser(Match); err != nil {
		// Jika terjadi error, kirim response dengan status 400 dan pesan error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	// Lanjutkan dengan proses penyimpanan ke database
	db := database.GetDB()
	result := db.Create(Match)

	// Jika terjadi error saat menyimpan, kirim response dengan status 500
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create match: " + result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    Match,
	})
}

// GetAllMatch godoc
// @Tags Match
// @Summary Get all Match
// @Description Get details of all match
// @ID get-all-Match
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Tbl_match
// @Security Bearer
// @Router /api/workflow/match [get]
func GetAllMatchs(c *fiber.Ctx) error {
	var Match []models.Tbl_match

	db := database.GetDB()
	result := db.Find(&Match)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Match,
	})
}

// @Tags Match
// @Summary Get match by ID
// @Description Get a match by ID
// @ID get-match-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Match ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Match tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/match/{id} [get]
func GetMatchById(c *fiber.Ctx) error {
	db := database.GetDB()

	MatchID := c.Params("id")

	var Match models.Tbl_match
	data := db.First(&Match, MatchID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Turnament tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Match,
	})
}

// @Tags Match
// @Summary Get match by ID
// @Description Get a match by ID
// @ID get-match-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Match ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Match tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/match/{id} [get]
func UpdateMatchById(c *fiber.Ctx) error {
	db := database.GetDB()

	MatchID := c.Params("id")

	var Match models.Tbl_match

	if err := db.First(&Match, MatchID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Match tidak ditemukan",
		})
	}
	updatedMatch := new(models.Tbl_match)
	if err := c.BodyParser(updatedMatch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	Match.Club_id1 = updatedMatch.Club_id1
	Match.Club_id2 = updatedMatch.Club_id2
	Match.Date_match = updatedMatch.Date_match
	Match.Level = updatedMatch.Level
	Match.Turnament_code = updatedMatch.Turnament_code

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&Match).Updates(&Match)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan Match",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    Match,
	})
}

// @Tags Match
// @Summary Delete match by ID
// @Description Delete a match by ID
// @ID delete-match-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Match ID"
// @Success 200 {object} map[string]interface{} "Berhasil Menghapus Data"
// @Failure 404 {object} map[string]interface{} "Match tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menghapus data Match"
// @Security Bearer
// @Router /api/workflow/match/{id} [delete]
func DeleteMatchById(c *fiber.Ctx) error {
	db := database.GetDB()

	MatchID := c.Params("id")

	var Match models.Tbl_match
	data := db.First(&Match, MatchID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Match tidak ditemukan",
		})
	}
	if err := db.Delete(&Match).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data Match",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data Match",
	})
}
