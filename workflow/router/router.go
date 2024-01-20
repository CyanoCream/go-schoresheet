package routerWorkflow

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/middleware"
	controllers "go-scoresheet/workflow/controller"
)

func InitializeRoutesWorkflow(api fiber.Router) {
	workflowGroup := api.Group("/workflow")
	workflowGroup.Use(middleware.Authentication) // Middleware yang spesifik untuk /workflow

	// Endpoint-endpoint untuk /turnament
	workflowGroup.Post("/turnament", controllers.CreateTurnament)
	workflowGroup.Get("/turnament", controllers.GetAllTurnaments)
	workflowGroup.Get("/turnament/:id", controllers.GetTurnamentById)
	workflowGroup.Post("/turnament/:id", controllers.UpdateTurnamentById)
	workflowGroup.Delete("/turnament/:id", controllers.DeleteTurnamentById)

	// Endpoint-endpoint untuk /score
	workflowGroup.Post("/score", controllers.CreateScore)
	workflowGroup.Get("/score", controllers.GetAllScores)
	workflowGroup.Get("/score/:id", controllers.GetScoreById)
	workflowGroup.Post("/score/:id", controllers.UpdateScoreById)
	workflowGroup.Delete("/score/:id", controllers.DeleteScoreById)

	// Endpoint-endpoint untuk /player
	workflowGroup.Post("/player", controllers.CreatePlayer)
	workflowGroup.Get("/player", controllers.GetAllPlayers)
	workflowGroup.Get("/player/:id", controllers.GetPlayerById)
	workflowGroup.Post("/player/:id", controllers.UpdatePlayerById)
	workflowGroup.Delete("/player/:id", controllers.DeletePlayerById)

	// Endpoint-endpoint untuk /match
	workflowGroup.Post("/match", controllers.CreateMatch)
	workflowGroup.Get("/match", controllers.GetAllMatchs)
	workflowGroup.Get("/match/:id", controllers.GetMatchById)
	workflowGroup.Post("/match/:id", controllers.UpdateMatchById)
	workflowGroup.Delete("/match/:id", controllers.DeleteMatchById)

	// Endpoint-endpoint untuk /club
	workflowGroup.Post("/club", controllers.CreateClub)
	workflowGroup.Get("/club", controllers.GetAllClubs)
	workflowGroup.Get("/club/:id", controllers.GetClubById)
	workflowGroup.Post("/club/:id", controllers.UpdateClubById)
	workflowGroup.Delete("/club/:id", controllers.DeleteClubById)

}
