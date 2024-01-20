// repository/player_repository.go
package repository

import (
	"go-scoresheet/workflow/models"
	"gorm.io/gorm"
)

type PlayerRepository interface {
	CreatePlayer(player *models.Tbl_player) error
	GetAllPlayers() ([]models.Tbl_player, error)
	GetPlayerById(id uint) (*models.Tbl_player, error)
	UpdatePlayer(player *models.Tbl_player) error
	DeletePlayer(id uint) error
}

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) PlayerRepository {
	return &playerRepository{
		db: db,
	}
}

func (r *playerRepository) CreatePlayer(player *models.Tbl_player) error {
	return r.db.Create(player).Error
}

func (r *playerRepository) GetAllPlayers() ([]models.Tbl_player, error) {
	var players []models.Tbl_player
	err := r.db.Find(&players).Error
	return players, err
}

func (r *playerRepository) GetPlayerById(id uint) (*models.Tbl_player, error) {
	var player models.Tbl_player
	err := r.db.First(&player, id).Error
	return &player, err
}

func (r *playerRepository) UpdatePlayer(player *models.Tbl_player) error {
	return r.db.Save(player).Error
}

func (r *playerRepository) DeletePlayer(id uint) error {
	return r.db.Delete(&models.Tbl_player{}, id).Error
}
