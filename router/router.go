package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "go-scoresheet/docs"
	"go-scoresheet/master/controllers"
	"go-scoresheet/master/router"
)

// @title GO-Scoresheet
// @version 2.0
// @description Documentation API GO-Scoresheet
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func InitializeRoutesMain() *fiber.App {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
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
	api.Post("/login", controllers.LoginUser)
	api.Delete("/logout", controllers.DeleteSessionByToken)
	routerMaster.InitializeRoutesMaster(api)
}
