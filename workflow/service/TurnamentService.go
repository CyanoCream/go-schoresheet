package service

import (
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/repository"
)

type TurnamentService interface {
	CreateTurnament(turnament *models.Turnament) error
	GetAllTurnaments() ([]models.Turnament, error)
	GetTurnamentById(id uint) (*models.Turnament, error)
	UpdateTurnament(turnament *models.Turnament) error
	DeleteTurnament(id uint) error
}

type turnamentService struct {
	turnamentRepository repository.TurnamentRepository
}

func NewTurnamentService(turnamentRepository repository.TurnamentRepository) TurnamentService {
	return &turnamentService{turnamentRepository}
}

func (s *turnamentService) CreateTurnament(turnament *models.Turnament) error {
	return s.turnamentRepository.CreateTurnament(turnament)
}

func (s *turnamentService) GetAllTurnaments() ([]models.Turnament, error) {
	return s.turnamentRepository.GetAllTurnaments()
}

func (s *turnamentService) GetTurnamentById(id uint) (*models.Turnament, error) {
	return s.turnamentRepository.GetTurnamentById(id)
}

func (s *turnamentService) UpdateTurnament(turnament *models.Turnament) error {
	return s.turnamentRepository.UpdateTurnament(turnament)
}

func (s *turnamentService) DeleteTurnament(id uint) error {
	return s.turnamentRepository.DeleteTurnament(id)
}
