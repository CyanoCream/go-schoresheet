// service/match_service.go
package service

import (
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/repository"
)

type MatchService interface {
	CreateMatch(match *models.Tbl_match) error
	GetAllMatches() ([]models.Tbl_match, error)
	GetMatchById(id uint) (*models.Tbl_match, error)
	UpdateMatch(match *models.Tbl_match) error
	DeleteMatch(id uint) error
}

type matchService struct {
	matchRepo repository.MatchRepository
}

func NewMatchService(matchRepo repository.MatchRepository) MatchService {
	return &matchService{
		matchRepo: matchRepo,
	}
}

func (s *matchService) CreateMatch(match *models.Tbl_match) error {
	return s.matchRepo.CreateMatch(match)
}

func (s *matchService) GetAllMatches() ([]models.Tbl_match, error) {
	return s.matchRepo.GetAllMatches()
}

func (s *matchService) GetMatchById(id uint) (*models.Tbl_match, error) {
	return s.matchRepo.GetMatchById(id)
}

func (s *matchService) UpdateMatch(match *models.Tbl_match) error {
	return s.matchRepo.UpdateMatch(match)
}

func (s *matchService) DeleteMatch(id uint) error {
	return s.matchRepo.DeleteMatch(id)
}
