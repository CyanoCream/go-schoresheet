package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/workflow/models"
)

// CreateTurnament godoc
// @Tags Turnament
// @Summary Create Turnament
// @Description Create New Turnament
// @ID createTurnament
// @Accept json
// @Produce json
// @Param requestBody body models.Turnament true "User credentials in JSON format"
// @Success 201 {object} models.Turnament
// @Security Bearer
// @Router /api/workflow/turnament [post]
func CreateTurnament(c *fiber.Ctx) error {
	Turnament := new(models.Turnament) // Pastikan ini adalah pointer ke struct yang benar

	// Parse body ke struct Turnament
	if err := c.BodyParser(Turnament); err != nil {
		// Jika terjadi error, kirim response dengan status 400 dan pesan error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	// Lanjutkan dengan proses penyimpanan ke database
	db := database.GetDB()
	result := db.Create(Turnament)

	// Jika terjadi error saat menyimpan, kirim response dengan status 500
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create turnament: " + result.Error.Error(),
		})
	}

	// Jika berhasil, kirim response dengan status 201 dan data Turnament yang telah dibuat
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    Turnament,
	})
}

// GetAllTurnament godoc
// @Tags Turnament
// @Summary Get all Turnament
// @Description Get details of all turnament
// @ID get-all-Turnament
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Turnament
// @Security Bearer
// @Router /api/workflow/turnament [get]
func GetAllTurnaments(c *fiber.Ctx) error {
	var Turnemant []models.Turnament

	db := database.GetDB()
	result := db.Find(&Turnemant)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Turnemant,
	})
}

// @Tags Turnament
// @Summary Get turnament by ID
// @Description Get a turnament by ID
// @ID get-turnament-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Turnament ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Turnament tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/turnament/{id} [get]
func GetTurnamentById(c *fiber.Ctx) error {
	db := database.GetDB()

	TurnamentID := c.Params("id")

	var Turnament models.Turnament
	data := db.First(&Turnament, TurnamentID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Turnament tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Turnament,
	})
}

// @Tags Turnament
// @Summary Get turnament by ID
// @Description Get a turnament by ID
// @ID get-turnament-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Turnament ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Turnament tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/turnament/{id} [get]
func UpdateTurnamentById(c *fiber.Ctx) error {
	db := database.GetDB()

	TurnamentID := c.Params("id")

	var Turnament models.Turnament

	if err := db.First(&Turnament, TurnamentID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Turnament tidak ditemukan",
		})
	}
	updatedTurnament := new(models.Turnament)
	if err := c.BodyParser(updatedTurnament); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	Turnament.Name = updatedTurnament.Name

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&Turnament).Updates(&Turnament)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan Permission",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    Turnament,
	})
}

// @Tags Turnament
// @Summary Delete turnament by ID
// @Description Delete a turnament by ID
// @ID delete-turnament-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Turnament ID"
// @Success 200 {object} map[string]interface{} "Berhasil Menghapus Data"
// @Failure 404 {object} map[string]interface{} "Turnament tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menghapus data Turnament"
// @Security Bearer
// @Router /api/workflow/turnament/{id} [delete]
func DeleteTurnamentById(c *fiber.Ctx) error {
	db := database.GetDB()

	TurnamentID := c.Params("id")

	var Turnament models.Turnament
	data := db.First(&Turnament, TurnamentID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Turnament tidak ditemukan",
		})
	}
	if err := db.Delete(&Turnament).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data Turnament",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data Turnament",
	})
}
