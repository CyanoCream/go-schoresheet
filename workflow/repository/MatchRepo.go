// repository/match_repository.go
package repository

import (
	"go-scoresheet/workflow/models"
	"gorm.io/gorm"
)

type MatchRepository interface {
	CreateMatch(match *models.Tbl_match) error
	GetAllMatches() ([]models.Tbl_match, error)
	GetMatchById(id uint) (*models.Tbl_match, error)
	UpdateMatch(match *models.Tbl_match) error
	DeleteMatch(id uint) error
}

type matchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) MatchRepository {
	return &matchRepository{
		db: db,
	}
}

func (r *matchRepository) CreateMatch(match *models.Tbl_match) error {
	return r.db.Create(match).Error
}

func (r *matchRepository) GetAllMatches() ([]models.Tbl_match, error) {
	var matches []models.Tbl_match
	err := r.db.Find(&matches).Error
	return matches, err
}

func (r *matchRepository) GetMatchById(id uint) (*models.Tbl_match, error) {
	var match models.Tbl_match
	err := r.db.First(&match, id).Error
	return &match, err
}

func (r *matchRepository) UpdateMatch(match *models.Tbl_match) error {
	return r.db.Save(match).Error
}

func (r *matchRepository) DeleteMatch(id uint) error {
	return r.db.Delete(&models.Tbl_match{}, id).Error
}
