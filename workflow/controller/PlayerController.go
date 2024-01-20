package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/database"
	"go-scoresheet/workflow/models"
)

// CreatePlayer godoc
// @Tags Player
// @Summary Create Player
// @Description Create New Player
// @ID createPlayer
// @Accept json
// @Produce json
// @Param requestBody body models.Tbl_player true "User credentials in JSON format"
// @Success 201 {object} models.Tbl_player
// @Security Bearer
// @Router /api/workflow/player [post]
func CreatePlayer(c *fiber.Ctx) error {
	Player := new(models.Tbl_player) // Pastikan ini adalah pointer ke struct yang benar

	if err := c.BodyParser(Player); err != nil {
		// Jika terjadi error, kirim response dengan status 400 dan pesan error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	// Lanjutkan dengan proses penyimpanan ke database
	db := database.GetDB()
	result := db.Create(Player)

	// Jika terjadi error saat menyimpan, kirim response dengan status 500
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create player: " + result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    Player,
	})
}

// GetAllPlayer godoc
// @Tags Player
// @Summary Get all Player
// @Description Get details of all player
// @ID get-all-Player
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Tbl_player
// @Security Bearer
// @Router /api/workflow/player [get]
func GetAllPlayers(c *fiber.Ctx) error {
	var Player []models.Tbl_player

	db := database.GetDB()
	result := db.Find(&Player)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch permissions",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Player,
	})
}

// @Tags Player
// @Summary Get player by ID
// @Description Get a player by ID
// @ID get-player-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Player ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Player tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/player/{id} [get]
func GetPlayerById(c *fiber.Ctx) error {
	db := database.GetDB()

	PlayerID := c.Params("id")

	var Player models.Tbl_player
	data := db.First(&Player, PlayerID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Turnament tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"data":    Player,
	})
}

// @Tags Player
// @Summary Get player by ID
// @Description Get a player by ID
// @ID get-player-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Player ID"
// @Success 200 {object} map[string]interface{} "success"
// @Failure 404 {object} map[string]interface{} "Player tidak ditemukan"
// @Security Bearer
// @Router /api/workflow/player/{id} [get]
func UpdatePlayerById(c *fiber.Ctx) error {
	db := database.GetDB()

	PlayerID := c.Params("id")

	var Player models.Tbl_player

	if err := db.First(&Player, PlayerID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Player tidak ditemukan",
		})
	}
	updatedPlayer := new(models.Tbl_player)
	if err := c.BodyParser(updatedPlayer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}
	Player.Name = updatedPlayer.Name
	Player.Club_id = updatedPlayer.Club_id
	Player.Position = updatedPlayer.Position

	// Menggunakan metode Updates untuk menyimpan perubahan ke database
	data := db.Model(&Player).Updates(&Player)

	if data.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal melakukan pembaruan Player",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil melakukan pembaruan",
		"data":    Player,
	})
}

// @Tags Player
// @Summary Delete player by ID
// @Description Delete a player by ID
// @ID delete-player-by-id
// @Accept  json
// @Produce  json
// @Param id path string true "Player ID"
// @Success 200 {object} map[string]interface{} "Berhasil Menghapus Data"
// @Failure 404 {object} map[string]interface{} "Player tidak ditemukan"
// @Failure 500 {object} map[string]interface{} "Gagal menghapus data Player"
// @Security Bearer
// @Router /api/workflow/player/{id} [delete]
func DeletePlayerById(c *fiber.Ctx) error {
	db := database.GetDB()

	PlayerID := c.Params("id")

	var Player models.Tbl_player
	data := db.First(&Player, PlayerID)

	if data.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data Player tidak ditemukan",
		})
	}
	if err := db.Delete(&Player).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data Player",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus Data Player",
	})
}
