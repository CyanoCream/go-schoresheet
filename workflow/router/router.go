package routerWorkflow

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/middleware"
	controllers "go-scoresheet/workflow/controller"
	WorkflowRepository "go-scoresheet/workflow/repository"
	WorkflowService "go-scoresheet/workflow/service"
	"gorm.io/gorm"
)

func InitializeRoutesWorkflow(api fiber.Router, db *gorm.DB) {
	workflowGroup := api.Group("/workflow")
	workflowGroup.Use(middleware.Authentication) // Middleware yang spesifik untuk /workflow

	// Endpoint-endpoint untuk /turnament
	turnamentController := controllers.NewTurnamentController(
		WorkflowService.NewTurnamentService(
			WorkflowRepository.NewTurnamentRepository(db),
		),
	)
	workflowGroup.Post("/turnament", turnamentController.CreateTurnament)
	workflowGroup.Get("/turnament", turnamentController.GetAllTurnaments)
	workflowGroup.Get("/turnament/:id", turnamentController.GetTurnamentById)
	workflowGroup.Post("/turnament/:id", turnamentController.UpdateTurnamentById)
	workflowGroup.Delete("/turnament/:id", turnamentController.DeleteTurnamentById)

	// Endpoint-endpoint untuk /score
	scoreController := controllers.NewScoreController(
		WorkflowService.NewScoreService(
			WorkflowRepository.NewScoreRepository(db),
		),
	)
	workflowGroup.Post("/score", scoreController.CreateScore)
	workflowGroup.Get("/score", scoreController.GetAllScores)
	workflowGroup.Get("/score/:id", scoreController.GetScoreById)
	workflowGroup.Post("/score/:id", scoreController.UpdateScoreById)
	workflowGroup.Delete("/score/:id", scoreController.DeleteScoreById)

	// Endpoint-endpoint untuk /player
	playerController := controllers.NewPlayerController(
		WorkflowService.NewPlayerService(
			WorkflowRepository.NewPlayerRepository(db),
		),
	)
	workflowGroup.Post("/player", playerController.CreatePlayer)
	workflowGroup.Get("/player", playerController.GetAllPlayers)
	workflowGroup.Get("/player/:id", playerController.GetPlayerById)
	workflowGroup.Post("/player/:id", playerController.UpdatePlayerById)
	workflowGroup.Delete("/player/:id", playerController.DeletePlayerById)

	// Endpoint-endpoint untuk /match
	matchController := controllers.NewMatchController(
		WorkflowService.NewMatchService(
			WorkflowRepository.NewMatchRepository(db),
		),
	)
	workflowGroup.Post("/match", matchController.CreateMatch)
	workflowGroup.Get("/match", matchController.GetAllMatches)
	workflowGroup.Get("/match/:id", matchController.GetMatchById)
	workflowGroup.Post("/match/:id", matchController.UpdateMatchById)
	workflowGroup.Delete("/match/:id", matchController.DeleteMatchById)

	// Endpoint-endpoint untuk /club
	clubController := controllers.NewClubController(
		WorkflowService.NewClubService(
			WorkflowRepository.NewClubRepository(db),
		),
	)
	workflowGroup.Post("/club", clubController.CreateClub)
	workflowGroup.Get("/club", clubController.GetAllClubs)
	workflowGroup.Get("/club/:id", clubController.GetClubById)
	workflowGroup.Post("/club/:id", clubController.UpdateClubById)
	workflowGroup.Delete("/club/:id", clubController.DeleteClubById)

}
