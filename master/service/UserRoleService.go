package service

import (
	"go-scoresheet/master/models"
	"go-scoresheet/master/repository"
)

type UserRoleService interface {
	CreateUserRole(userRole *models.UserRole) error
	GetAllUserRoles() ([]models.UserRole, error)
	GetUserRoleById(id uint) (*models.UserRole, error)
	UpdateUserRole(userRole *models.UserRole) error
	DeleteUserRole(id uint) error
}

type userRoleService struct {
	userRoleRepo repository.UserRoleRepository
}

func NewUserRoleService(userRoleRepo repository.UserRoleRepository) UserRoleService {
	return &userRoleService{
		userRoleRepo: userRoleRepo,
	}
}

func (s *userRoleService) CreateUserRole(userRole *models.UserRole) error {
	return s.userRoleRepo.CreateUserRole(userRole)
}

func (s *userRoleService) GetAllUserRoles() ([]models.UserRole, error) {
	return s.userRoleRepo.GetAllUserRoles()
}

func (s *userRoleService) GetUserRoleById(id uint) (*models.UserRole, error) {
	return s.userRoleRepo.GetUserRoleById(id)
}

func (s *userRoleService) UpdateUserRole(userRole *models.UserRole) error {
	return s.userRoleRepo.UpdateUserRole(userRole)
}

func (s *userRoleService) DeleteUserRole(id uint) error {
	return s.userRoleRepo.DeleteUserRole(id)
}
