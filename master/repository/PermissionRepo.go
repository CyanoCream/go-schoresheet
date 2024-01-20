// repository/permission_repository.go
package repository

import (
	"go-scoresheet/master/models"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	CreatePermission(permission *models.Permission) error
	GetAllPermissions() ([]models.Permission, error)
	GetPermissionById(id uint) (*models.Permission, error)
	UpdatePermission(permission *models.Permission) error
	DeletePermission(id uint) error
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository() PermissionRepository {
	var db *gorm.DB
	return &permissionRepository{
		db: db,
	}
}

func (r *permissionRepository) CreatePermission(permission *models.Permission) error {
	return r.db.Create(permission).Error
}

func (r *permissionRepository) GetAllPermissions() ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) GetPermissionById(id uint) (*models.Permission, error) {
	var permission models.Permission
	err := r.db.First(&permission, id).Error
	return &permission, err
}

func (r *permissionRepository) UpdatePermission(permission *models.Permission) error {
	return r.db.Save(permission).Error
}

func (r *permissionRepository) DeletePermission(id uint) error {
	return r.db.Delete(&models.Permission{}, id).Error
}
