package router

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/master/controllers"
	"go-scoresheet/master/router"
)

func InitializeRoutesMain() *fiber.App {
	app := fiber.New()

	SetupRoutes(app)
	// Rute untuk membuat pengguna baru
	//app.Post("/users", controllers.CreateUser)
	//
	//// Rute untuk mendapatkan semua pengguna
	//app.Get("/users", controllers.GetAllUsers)

	return app
}

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/users/login", controllers.LoginUser)
	routerMaster.InitializeRoutesMaster(api)
}
