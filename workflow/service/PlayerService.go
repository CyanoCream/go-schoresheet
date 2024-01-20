// service/player_service.go
package service

import (
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/repository"
)

type PlayerService interface {
	CreatePlayer(player *models.Tbl_player) error
	GetAllPlayers() ([]models.Tbl_player, error)
	GetPlayerById(id uint) (*models.Tbl_player, error)
	UpdatePlayer(player *models.Tbl_player) error
	DeletePlayer(id uint) error
}

type playerService struct {
	playerRepo repository.PlayerRepository
}

func NewPlayerService(playerRepo repository.PlayerRepository) PlayerService {
	return &playerService{
		playerRepo: playerRepo,
	}
}

func (s *playerService) CreatePlayer(player *models.Tbl_player) error {
	return s.playerRepo.CreatePlayer(player)
}

func (s *playerService) GetAllPlayers() ([]models.Tbl_player, error) {
	return s.playerRepo.GetAllPlayers()
}

func (s *playerService) GetPlayerById(id uint) (*models.Tbl_player, error) {
	return s.playerRepo.GetPlayerById(id)
}

func (s *playerService) UpdatePlayer(player *models.Tbl_player) error {
	return s.playerRepo.UpdatePlayer(player)
}

func (s *playerService) DeletePlayer(id uint) error {
	return s.playerRepo.DeletePlayer(id)
}
