package routerMaster

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/master/controllers"
	"go-scoresheet/middleware"
)

func InitializeRoutesMaster(api fiber.Router) {
	api.Post("/users", controllers.CreateUser)
	api.Use(middleware.Authentication)

	//Master User

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

	//Master User Role
	api.Post("/user-role", controllers.CreateUserRole)
	api.Get("/user-role", controllers.GetAllUserRoles)
	api.Get("/user-role/:id", controllers.GetUserRoleById)
	api.Post("/user-role/:id", controllers.UpdateUserRoleById)
	api.Delete("/user-role/:id", controllers.DeleteUserRoleById)

	//Master Permission Role
	api.Post("/permission-role", controllers.CreatePermissionRole)
	api.Get("/permission-role", controllers.GetAllPermissionRoles)
	api.Get("/permission-role/:id", controllers.GetPermissionRoleById)
	api.Post("/permission-role/:id", controllers.UpdatePermissionRoleById)
	api.Delete("/permission-role/:id", controllers.DeletePermissionRoleById)

}
