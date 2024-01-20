package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/service"
	"strconv"
)

type ClubController struct {
	clubService service.ClubService
}

func NewClubController(clubService service.ClubService) *ClubController {
	return &ClubController{
		clubService: clubService,
	}
}

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
func (c *ClubController) CreateClub(ctx *fiber.Ctx) error {
	club := new(models.Tbl_club)

	if err := ctx.BodyParser(club); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON: " + err.Error(),
		})
	}

	if err := c.clubService.CreateClub(club); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create club: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    club,
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
func (c *ClubController) GetAllClubs(ctx *fiber.Ctx) error {
	clubs, err := c.clubService.GetAllClubs()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch clubs",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    clubs,
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
func (c *ClubController) GetClubById(ctx *fiber.Ctx) error {
	clubID := ctx.Params("id")
	id, err := strconv.ParseUint(clubID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Club ID",
		})
	}

	club, err := c.clubService.GetClubById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Club not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    club,
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
func (c *ClubController) UpdateClubById(ctx *fiber.Ctx) error {
	clubID := ctx.Params("id")
	id, err := strconv.ParseUint(clubID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Club ID",
		})
	}

	club, err := c.clubService.GetClubById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Club not found",
		})
	}

	updatedClub := new(models.Tbl_club)
	if err := ctx.BodyParser(updatedClub); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON",
		})
	}

	club.Name = updatedClub.Name
	club.Hometown = updatedClub.Hometown

	if err := c.clubService.UpdateClub(club); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Club",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully updated Club",
		"data":    club,
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
func (c *ClubController) DeleteClubById(ctx *fiber.Ctx) error {
	clubID := ctx.Params("id")
	id, err := strconv.ParseUint(clubID, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Club ID",
		})
	}

	if err := c.clubService.DeleteClub(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete Club",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Successfully deleted Club",
	})
}
