package routerMaster

import (
	"github.com/gofiber/fiber/v2"
	"go-scoresheet/master/controllers"
	MasterRepository "go-scoresheet/master/repository"
	MasterService "go-scoresheet/master/service"
	"go-scoresheet/middleware"
)

func InitializeRoutesMaster(api fiber.Router) {
	UserController := controllers.NewUserController(
		MasterService.NewUserService(
			MasterRepository.NewUserRepository(),
		),
	)
	api.Post("/users", UserController.CreateUser)
	api.Get("/users", UserController.GetAllUsers)
	api.Use(middleware.Authentication)

	//Master User

	api.Get("/users/:id", UserController.GetUserById)
	api.Post("/users/:id", UserController.UpdateUserById)
	api.Delete("/users/:id", UserController.DeleteUserById)

	//Master Permission
	PermissionController := controllers.NewPermissionController(
		MasterService.NewPermissionService(
			MasterRepository.NewPermissionRepository(),
		),
	)
	api.Post("/permission", PermissionController.CreatePermission)
	api.Get("/permission", PermissionController.GetAllPermissions)
	api.Get("/permission/:id", PermissionController.GetPermissionById)
	api.Post("/permission/:id", PermissionController.UpdatePermissionById)
	api.Delete("/permission/:id", PermissionController.DeletePermissionById)

	//Master Role
	RoleController := controllers.NewRoleController(
		MasterService.NewRoleService(
			MasterRepository.NewRoleRepository(),
		),
	)
	api.Post("/role", RoleController.CreateRole)
	api.Get("/role", RoleController.GetAllRoles)
	api.Get("/role/:id", RoleController.GetRoleById)
	api.Post("/role/:id", RoleController.UpdateRoleById)
	api.Delete("/role/:id", RoleController.DeleteRoleById)

	//Master User Role
	userRoleController := controllers.NewUserRoleController(
		MasterService.NewUserRoleService(
			MasterRepository.NewUserRoleRepository(),
		),
	)
	api.Post("/user-role", userRoleController.CreateUserRole)
	api.Get("/user-role", userRoleController.GetAllUserRoles)
	api.Get("/user-role/:id", userRoleController.GetUserRoleById)
	api.Post("/user-role/:id", userRoleController.UpdateUserRoleById)
	api.Delete("/user-role/:id", userRoleController.DeleteUserRoleById)

	//Master Permission Role
	PermissionRoleController := controllers.NewPermissionRoleController(
		MasterService.NewPermissionRoleService(
			MasterRepository.NewPermissionRoleRepository(),
		),
	)
	api.Post("/permission-role", PermissionRoleController.CreatePermissionRole)
	api.Get("/permission-role", PermissionRoleController.GetAllPermissionRoles)
	api.Get("/permission-role/:id", PermissionRoleController.GetPermissionRoleById)
	api.Post("/permission-role/:id", PermissionRoleController.UpdatePermissionRoleById)
	api.Delete("/permission-role/:id", PermissionRoleController.DeletePermissionRoleById)

}
