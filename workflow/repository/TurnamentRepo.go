// repository/turnament_repository.go
package repository

import (
	"go-scoresheet/workflow/models"
	"gorm.io/gorm"
)

type TurnamentRepository interface {
	CreateTurnament(turnament *models.Turnament) error
	GetAllTurnaments() ([]models.Turnament, error)
	GetTurnamentById(id uint) (*models.Turnament, error)
	UpdateTurnament(turnament *models.Turnament) error
	DeleteTurnament(id uint) error
}

type turnamentRepository struct {
	db *gorm.DB
}

func NewTurnamentRepository(db *gorm.DB) TurnamentRepository {
	return &turnamentRepository{db}
}

func (r *turnamentRepository) CreateTurnament(turnament *models.Turnament) error {
	return r.db.Create(turnament).Error
}

func (r *turnamentRepository) GetAllTurnaments() ([]models.Turnament, error) {
	var turnaments []models.Turnament
	if err := r.db.Find(&turnaments).Error; err != nil {
		return nil, err
	}
	return turnaments, nil
}

func (r *turnamentRepository) GetTurnamentById(id uint) (*models.Turnament, error) {
	var turnament models.Turnament
	if err := r.db.First(&turnament, id).Error; err != nil {
		return nil, err
	}
	return &turnament, nil
}

func (r *turnamentRepository) UpdateTurnament(turnament *models.Turnament) error {
	return r.db.Save(turnament).Error
}

func (r *turnamentRepository) DeleteTurnament(id uint) error {
	return r.db.Delete(&models.Turnament{}, id).Error
}
