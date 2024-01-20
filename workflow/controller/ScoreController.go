package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/service"
	"strconv"
)

type ScoreController struct {
	scoreService service.ScoreService
}

func NewScoreController(scoreService service.ScoreService) *ScoreController {
	return &ScoreController{
		scoreService: scoreService,
	}
}

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
func (c *ScoreController) CreateScore(ctx *fiber.Ctx) error {
	score := new(models.Tbl_score)

	if err := ctx.BodyParser(score); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	if err := c.scoreService.CreateScore(score); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create score: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    score,
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
func (c *ScoreController) GetAllScores(ctx *fiber.Ctx) error {
	scores, err := c.scoreService.GetAllScores()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch scores",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    scores,
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
func (c *ScoreController) GetScoreById(ctx *fiber.Ctx) error {
	scoreID := ctx.Params("id")
	id, err := strconv.ParseUint(scoreID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Score ID",
		})
	}

	score, err := c.scoreService.GetScoreById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Score not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    score,
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
func (c *ScoreController) UpdateScoreById(ctx *fiber.Ctx) error {
	scoreID := ctx.Params("id")
	id, err := strconv.ParseUint(scoreID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Score ID",
		})
	}

	score, err := c.scoreService.GetScoreById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Score not found",
		})
	}

	updatedScore := new(models.Tbl_score)
	if err := ctx.BodyParser(updatedScore); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	score.Match_id = updatedScore.Match_id
	score.Club_id = updatedScore.Club_id
	score.Score = updatedScore.Score
	score.Status = updatedScore.Status

	if err := c.scoreService.UpdateScore(score); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Score",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated Score",
		"data":    score,
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
func (c *ScoreController) DeleteScoreById(ctx *fiber.Ctx) error {
	scoreID := ctx.Params("id")
	id, err := strconv.ParseUint(scoreID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Score ID",
		})
	}

	if err := c.scoreService.DeleteScore(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Score",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted Score",
	})
}
