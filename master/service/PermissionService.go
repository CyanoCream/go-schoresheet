// service/permission_service.go
package service

import (
	"go-scoresheet/master/models"
	"go-scoresheet/master/repository"
)

type PermissionService interface {
	CreatePermission(permission *models.Permission) error
	GetAllPermissions() ([]models.Permission, error)
	GetPermissionById(id uint) (*models.Permission, error)
	UpdatePermission(permission *models.Permission) error
	DeletePermission(id uint) error
}

type permissionService struct {
	permissionRepo repository.PermissionRepository
}

func NewPermissionService(permissionRepo repository.PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepo: permissionRepo,
	}
}

func (s *permissionService) CreatePermission(permission *models.Permission) error {
	return s.permissionRepo.CreatePermission(permission)
}

func (s *permissionService) GetAllPermissions() ([]models.Permission, error) {
	return s.permissionRepo.GetAllPermissions()
}

func (s *permissionService) GetPermissionById(id uint) (*models.Permission, error) {
	return s.permissionRepo.GetPermissionById(id)
}

func (s *permissionService) UpdatePermission(permission *models.Permission) error {
	return s.permissionRepo.UpdatePermission(permission)
}

func (s *permissionService) DeletePermission(id uint) error {
	return s.permissionRepo.DeletePermission(id)
}
