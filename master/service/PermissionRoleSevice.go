// service/permission_role_service.go
package service

import (
	"go-scoresheet/master/models"
	"go-scoresheet/master/repository"
)

type PermissionRoleService interface {
	CreatePermissionRole(permissionRole *models.PermissionRole) error
	GetAllPermissionRoles() ([]models.PermissionRole, error)
	GetPermissionRoleById(id uint) (*models.PermissionRole, error)
	UpdatePermissionRole(permissionRole *models.PermissionRole) error
	DeletePermissionRole(id uint) error
}

type permissionRoleService struct {
	permissionRoleRepo repository.PermissionRoleRepository
}

func NewPermissionRoleService(permissionRoleRepo repository.PermissionRoleRepository) PermissionRoleService {
	return &permissionRoleService{
		permissionRoleRepo: permissionRoleRepo,
	}
}

func (s *permissionRoleService) CreatePermissionRole(permissionRole *models.PermissionRole) error {
	return s.permissionRoleRepo.CreatePermissionRole(permissionRole)
}

func (s *permissionRoleService) GetAllPermissionRoles() ([]models.PermissionRole, error) {
	return s.permissionRoleRepo.GetAllPermissionRoles()
}

func (s *permissionRoleService) GetPermissionRoleById(id uint) (*models.PermissionRole, error) {
	return s.permissionRoleRepo.GetPermissionRoleById(id)
}

func (s *permissionRoleService) UpdatePermissionRole(permissionRole *models.PermissionRole) error {
	return s.permissionRoleRepo.UpdatePermissionRole(permissionRole)
}

func (s *permissionRoleService) DeletePermissionRole(id uint) error {
	return s.permissionRoleRepo.DeletePermissionRole(id)
}
