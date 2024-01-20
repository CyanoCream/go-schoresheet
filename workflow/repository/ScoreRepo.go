// repository/score_repository.go
package repository

import (
	"go-scoresheet/workflow/models"
	"gorm.io/gorm"
)

type ScoreRepository interface {
	CreateScore(score *models.Tbl_score) error
	GetAllScores() ([]models.Tbl_score, error)
	GetScoreById(id uint) (*models.Tbl_score, error)
	UpdateScore(score *models.Tbl_score) error
	DeleteScore(id uint) error
}

type scoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) ScoreRepository {
	return &scoreRepository{
		db: db,
	}
}

func (r *scoreRepository) CreateScore(score *models.Tbl_score) error {
	return r.db.Create(score).Error
}

func (r *scoreRepository) GetAllScores() ([]models.Tbl_score, error) {
	var scores []models.Tbl_score
	err := r.db.Find(&scores).Error
	return scores, err
}

func (r *scoreRepository) GetScoreById(id uint) (*models.Tbl_score, error) {
	var score models.Tbl_score
	err := r.db.First(&score, id).Error
	return &score, err
}

func (r *scoreRepository) UpdateScore(score *models.Tbl_score) error {
	return r.db.Save(score).Error
}

func (r *scoreRepository) DeleteScore(id uint) error {
	return r.db.Delete(&models.Tbl_score{}, id).Error
}
