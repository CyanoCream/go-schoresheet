// service/role_service.go
package service

import (
	"go-scoresheet/master/models"
	"go-scoresheet/master/repository"
)

type RoleService interface {
	CreateRole(role *models.Role) error
	GetAllRoles() ([]models.Role, error)
	GetRoleById(id uint) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

func (s *roleService) CreateRole(role *models.Role) error {
	return s.roleRepo.CreateRole(role)
}

func (s *roleService) GetAllRoles() ([]models.Role, error) {
	return s.roleRepo.GetAllRoles()
}

func (s *roleService) GetRoleById(id uint) (*models.Role, error) {
	return s.roleRepo.GetRoleById(id)
}

func (s *roleService) UpdateRole(role *models.Role) error {
	return s.roleRepo.UpdateRole(role)
}

func (s *roleService) DeleteRole(id uint) error {
	return s.roleRepo.DeleteRole(id)
}
