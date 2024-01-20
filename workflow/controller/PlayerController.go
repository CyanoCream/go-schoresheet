package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/service"
	"strconv"
)

type PlayerController struct {
	playerService service.PlayerService
}

func NewPlayerController(playerService service.PlayerService) *PlayerController {
	return &PlayerController{
		playerService: playerService,
	}
}

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
func (c *PlayerController) CreatePlayer(ctx *fiber.Ctx) error {
	player := new(models.Tbl_player)

	if err := ctx.BodyParser(player); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	if err := c.playerService.CreatePlayer(player); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create player: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    player,
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
func (c *PlayerController) GetAllPlayers(ctx *fiber.Ctx) error {
	players, err := c.playerService.GetAllPlayers()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch players",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    players,
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
func (c *PlayerController) GetPlayerById(ctx *fiber.Ctx) error {
	playerID := ctx.Params("id")
	id, err := strconv.ParseUint(playerID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Player ID",
		})
	}

	player, err := c.playerService.GetPlayerById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Player not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    player,
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
func (c *PlayerController) UpdatePlayerById(ctx *fiber.Ctx) error {
	playerID := ctx.Params("id")
	id, err := strconv.ParseUint(playerID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Player ID",
		})
	}

	player, err := c.playerService.GetPlayerById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Player not found",
		})
	}

	updatedPlayer := new(models.Tbl_player)
	if err := ctx.BodyParser(updatedPlayer); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	player.Name = updatedPlayer.Name
	player.Club_id = updatedPlayer.Club_id
	player.Position = updatedPlayer.Position

	if err := c.playerService.UpdatePlayer(player); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Player",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated Player",
		"data":    player,
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
func (c *PlayerController) DeletePlayerById(ctx *fiber.Ctx) error {
	playerID := ctx.Params("id")
	id, err := strconv.ParseUint(playerID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Player ID",
		})
	}

	if err := c.playerService.DeletePlayer(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Player",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted Player",
	})
}
