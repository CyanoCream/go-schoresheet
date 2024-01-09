package routerMaster

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/master/controllers"
	"go-scoresheet/middleware"
)

func InitializeRoutesMaster(api fiber.Router) {
	api.Use(middleware.Authentication)
	api.Delete("/logout", controllers.DeleteSessionByToken)
	//Master User
	api.Post("/users", controllers.CreateUser)
	api.Get("/users", controllers.GetAllUsers)
	api.Get("/users/:id", controllers.GetUserById)
	api.Post("/users/:id", controllers.UpdateUserById)
	api.Delete("/users/:id", controllers.DeleteUserById)

	//Master Permission
	api.Post("/permission", controllers.CreatePermission)
	api.Get("/permission", controllers.GetAllPermissions)
	api.Get("/permission/:id", controllers.GetPermissionById)
	api.Post("/permission/:id", controllers.UpdatePermissionById)
	api.Delete("/permission/:id", controllers.DeletePermissionById)

	//Master Role
	api.Post("/role", controllers.CreateRole)
	api.Get("/role", controllers.GetAllRoles)
	api.Get("/role/:id", controllers.GetRoleById)
	api.Post("/role/:id", controllers.UpdateRoleById)
	api.Delete("/role/:id", controllers.DeleteRoleById)

	//Master Permission Role
	api.Post("/permission-role", controllers.CreatePermissionRole)
	api.Get("/permission-role", controllers.GetAllPermissionRoles)
	api.Get("/permission-role/:id", controllers.GetPermissionRoleById)
	api.Post("/permission-role/:id", controllers.UpdateRoleById)
	api.Delete("/permission-role/:id", controllers.DeleteRoleById)

}
