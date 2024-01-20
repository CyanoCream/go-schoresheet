package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/service"
	"strconv"
)

type MatchController struct {
	matchService service.MatchService
}

func NewMatchController(matchService service.MatchService) *MatchController {
	return &MatchController{
		matchService: matchService,
	}
}

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
func (c *MatchController) CreateMatch(ctx *fiber.Ctx) error {
	match := new(models.Tbl_match)

	if err := ctx.BodyParser(match); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	if err := c.matchService.CreateMatch(match); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create match: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    match,
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
func (c *MatchController) GetAllMatches(ctx *fiber.Ctx) error {
	matches, err := c.matchService.GetAllMatches()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch matches",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    matches,
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
func (c *MatchController) GetMatchById(ctx *fiber.Ctx) error {
	matchID := ctx.Params("id")
	id, err := strconv.ParseUint(matchID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Match ID",
		})
	}

	match, err := c.matchService.GetMatchById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Match not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    match,
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
func (c *MatchController) UpdateMatchById(ctx *fiber.Ctx) error {
	matchID := ctx.Params("id")
	id, err := strconv.ParseUint(matchID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Match ID",
		})
	}

	match, err := c.matchService.GetMatchById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Match not found",
		})
	}

	updatedMatch := new(models.Tbl_match)
	if err := ctx.BodyParser(updatedMatch); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	match.Club_id1 = updatedMatch.Club_id1
	match.Club_id2 = updatedMatch.Club_id2
	match.Date_match = updatedMatch.Date_match
	match.Level = updatedMatch.Level
	match.Turnament_code = updatedMatch.Turnament_code

	if err := c.matchService.UpdateMatch(match); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Match",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated Match",
		"data":    match,
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
func (c *MatchController) DeleteMatchById(ctx *fiber.Ctx) error {
	matchID := ctx.Params("id")
	id, err := strconv.ParseUint(matchID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Match ID",
		})
	}

	if err := c.matchService.DeleteMatch(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Match",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted Match",
	})
}
