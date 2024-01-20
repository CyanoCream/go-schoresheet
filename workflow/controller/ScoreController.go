package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/workflow/models"
)

// CreateScore godoc
// @Tags Score
// @Summary Create Score
// @Description Create New Score
// @ID createScore
// @Accept json
// @Produce json
// @Param requestBody body models.Tbl_score true "User credentials in JSON format"
// @Success 201 {object} models.Tbl_score
// @Security Bearer
// @Router /api/workflow/score [post]
func CreateScore(c *fiber.Ctx) error {
	Score := new(models.Tbl_score) // Pastikan ini adalah pointer ke struct yang benar

	// Parse body ke struct Score
	if err := c.BodyParser(Score); err != nil {
		// Jika terjadi error, kirim response dengan status 400 dan pesan error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	// Lanjutkan dengan proses penyimpanan ke database
	db := database.GetDB()
	result := db.Create(Score)

	// Jika terjadi error saat menyimpan, kirim response dengan status 500
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create score: " + result.Error.Error(),
		})
	}

	// Jika berhasil, kirim response dengan status 201 dan data Score yang telah dibuat
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    Score,
	})
}

// GetAllScore godoc
// @Tags Score
// @Summary Get all Score
// @Description Get details of all score
// @ID get-all-Score
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Tbl_score
// @Security Bearer
// @Router /api/workflow/score [get]
func GetAllScores(c *fiber.Ctx) error {
	var Score []models.Tbl_score

	db := database.GetDB()
	result := db.Find(&Score)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Score,
	})
}

// @Tags Score
// @Summary Get score by ID
// @Description Get a score by ID
// @ID get-score-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Score ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Score tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/score/{id} [get]
func GetScoreById(c *fiber.Ctx) error {
	db := database.GetDB()

	ScoreID := c.Params("id")

	var Score models.Tbl_score
	data := db.First(&Score, ScoreID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Turnament tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Score,
	})
}

// @Tags Score
// @Summary Get score by ID
// @Description Get a score by ID
// @ID get-score-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Score ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Score tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/score/{id} [get]
func UpdateScoreById(c *fiber.Ctx) error {
	db := database.GetDB()

	ScoreID := c.Params("id")

	var Score models.Tbl_score

	if err := db.First(&Score, ScoreID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Score tidak ditemukan",
		})
	}
	updatedScore := new(models.Tbl_score)
	if err := c.BodyParser(updatedScore); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	Score.Match_id = updatedScore.Match_id
	Score.Club_id = updatedScore.Club_id
	Score.Score = updatedScore.Score
	Score.Status = updatedScore.Status

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&Score).Updates(&Score)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan Permission",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    Score,
	})
}

// @Tags Score
// @Summary Delete score by ID
// @Description Delete a score by ID
// @ID delete-score-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Score ID"
// @Success 200 {object} map[string]interface{} "Berhasil Menghapus Data"
// @Failure 404 {object} map[string]interface{} "Score tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menghapus data Score"
// @Security Bearer
// @Router /api/workflow/score/{id} [delete]
func DeleteScoreById(c *fiber.Ctx) error {
	db := database.GetDB()

	ScoreID := c.Params("id")

	var Score models.Tbl_score
	data := db.First(&Score, ScoreID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Score tidak ditemukan",
		})
	}
	if err := db.Delete(&Score).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data Score",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data Score",
	})
}
