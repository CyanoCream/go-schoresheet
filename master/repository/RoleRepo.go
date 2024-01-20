// repository/role_repository.go
package repository

import (
	"go-scoresheet/master/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *models.Role) error
	GetAllRoles() ([]models.Role, error)
	GetRoleById(id uint) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository() RoleRepository {
	var db *gorm.DB
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) CreateRole(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *roleRepository) GetRoleById(id uint) (*models.Role, error) {
	var role models.Role
	err := r.db.First(&role, id).Error
	return &role, err
}

func (r *roleRepository) UpdateRole(role *models.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) DeleteRole(id uint) error {
	return r.db.Delete(&models.Role{}, id).Error
}
