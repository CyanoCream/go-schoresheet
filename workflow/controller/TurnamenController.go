package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/service"
	"strconv"
)

type TurnamentController struct {
	turnamentService service.TurnamentService
}

func NewTurnamentController(turnamentService service.TurnamentService) TurnamentController {
	return TurnamentController{turnamentService}
}

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

func (c *TurnamentController) CreateTurnament(ctx *fiber.Ctx) error {
	turnament := new(models.Turnament)

	if err := ctx.BodyParser(turnament); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	if err := c.turnamentService.CreateTurnament(turnament); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create turnament: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    turnament,
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
func (c *TurnamentController) GetAllTurnaments(ctx *fiber.Ctx) error {
	turnaments, err := c.turnamentService.GetAllTurnaments()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch turnaments",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    turnaments,
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
func (c *TurnamentController) GetTurnamentById(ctx *fiber.Ctx) error {
	turnamentID := ctx.Params("id")
	id, err := strconv.ParseUint(turnamentID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Turnament ID",
		})
	}

	turnament, err := c.turnamentService.GetTurnamentById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Turnament not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    turnament,
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
func (c *TurnamentController) UpdateTurnamentById(ctx *fiber.Ctx) error {
	turnamentID := ctx.Params("id")
	id, err := strconv.ParseUint(turnamentID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Turnament ID",
		})
	}

	turnament, err := c.turnamentService.GetTurnamentById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Turnament not found",
		})
	}

	updatedTurnament := new(models.Turnament)
	if err := ctx.BodyParser(updatedTurnament); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	turnament.Name = updatedTurnament.Name

	if err := c.turnamentService.UpdateTurnament(turnament); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Turnament",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated Turnament",
		"data":    turnament,
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
func (c *TurnamentController) DeleteTurnamentById(ctx *fiber.Ctx) error {
	turnamentID := ctx.Params("id")
	id, err := strconv.ParseUint(turnamentID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Turnament ID",
		})
	}

	if err := c.turnamentService.DeleteTurnament(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Turnament",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted Turnament",
	})
}
