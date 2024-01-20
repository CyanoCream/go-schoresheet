// service/score_service.go
package service

import (
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/repository"
)

type ScoreService interface {
	CreateScore(score *models.Tbl_score) error
	GetAllScores() ([]models.Tbl_score, error)
	GetScoreById(id uint) (*models.Tbl_score, error)
	UpdateScore(score *models.Tbl_score) error
	DeleteScore(id uint) error
}

type scoreService struct {
	scoreRepo repository.ScoreRepository
}

func NewScoreService(scoreRepo repository.ScoreRepository) ScoreService {
	return &scoreService{
		scoreRepo: scoreRepo,
	}
}

func (s *scoreService) CreateScore(score *models.Tbl_score) error {
	return s.scoreRepo.CreateScore(score)
}

func (s *scoreService) GetAllScores() ([]models.Tbl_score, error) {
	return s.scoreRepo.GetAllScores()
}

func (s *scoreService) GetScoreById(id uint) (*models.Tbl_score, error) {
	return s.scoreRepo.GetScoreById(id)
}

func (s *scoreService) UpdateScore(score *models.Tbl_score) error {
	return s.scoreRepo.UpdateScore(score)
}

func (s *scoreService) DeleteScore(id uint) error {
	return s.scoreRepo.DeleteScore(id)
}
