// repository/user_role_repository.go
package repository

import (
	"go-scoresheet/master/models"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	CreateUserRole(userRole *models.UserRole) error
	GetAllUserRoles() ([]models.UserRole, error)
	GetUserRoleById(id uint) (*models.UserRole, error)
	UpdateUserRole(userRole *models.UserRole) error
	DeleteUserRole(id uint) error
}

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository() UserRoleRepository {
	var db *gorm.DB
	return &userRoleRepository{
		db: db,
	}
}

func (r *userRoleRepository) CreateUserRole(userRole *models.UserRole) error {
	return r.db.Create(userRole).Error
}

func (r *userRoleRepository) GetAllUserRoles() ([]models.UserRole, error) {
	var userRoles []models.UserRole
	err := r.db.Find(&userRoles).Error
	return userRoles, err
}

func (r *userRoleRepository) GetUserRoleById(id uint) (*models.UserRole, error) {
	var userRole models.UserRole
	err := r.db.First(&userRole, id).Error
	return &userRole, err
}

func (r *userRoleRepository) UpdateUserRole(userRole *models.UserRole) error {
	return r.db.Save(userRole).Error
}

func (r *userRoleRepository) DeleteUserRole(id uint) error {
	return r.db.Delete(&models.UserRole{}, id).Error
}
