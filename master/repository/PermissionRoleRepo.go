// repository/permission_role_repository.go
package repository

import (
	"go-scoresheet/master/models"
	"gorm.io/gorm"
)

type PermissionRoleRepository interface {
	CreatePermissionRole(permissionRole *models.PermissionRole) error
	GetAllPermissionRoles() ([]models.PermissionRole, error)
	GetPermissionRoleById(id uint) (*models.PermissionRole, error)
	UpdatePermissionRole(permissionRole *models.PermissionRole) error
	DeletePermissionRole(id uint) error
}

type permissionRoleRepository struct {
	db *gorm.DB
}

func NewPermissionRoleRepository(db *gorm.DB) PermissionRoleRepository {
	return &permissionRoleRepository{
		db: db,
	}
}

func (r *permissionRoleRepository) CreatePermissionRole(permissionRole *models.PermissionRole) error {
	return r.db.Create(permissionRole).Error
}

func (r *permissionRoleRepository) GetAllPermissionRoles() ([]models.PermissionRole, error) {
	var permissionRoles []models.PermissionRole
	err := r.db.Find(&permissionRoles).Error
	return permissionRoles, err
}

func (r *permissionRoleRepository) GetPermissionRoleById(id uint) (*models.PermissionRole, error) {
	var permissionRole models.PermissionRole
	err := r.db.First(&permissionRole, id).Error
	return &permissionRole, err
}

func (r *permissionRoleRepository) UpdatePermissionRole(permissionRole *models.PermissionRole) error {
	return r.db.Save(permissionRole).Error
}

func (r *permissionRoleRepository) DeletePermissionRole(id uint) error {
	return r.db.Delete(&models.PermissionRole{}, id).Error
}
