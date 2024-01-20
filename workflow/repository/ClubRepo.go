// repository/club_repository.go
package repository

import (
	"go-scoresheet/workflow/models"
	"gorm.io/gorm"
)

type ClubRepository interface {
	CreateClub(club *models.Tbl_club) error
	GetAllClubs() ([]models.Tbl_club, error)
	GetClubById(id uint) (*models.Tbl_club, error)
	UpdateClub(club *models.Tbl_club) error
	DeleteClub(id uint) error
}

type clubRepository struct {
	db *gorm.DB
}

func NewClubRepository(db *gorm.DB) ClubRepository {
	return &clubRepository{
		db: db,
	}
}

func (r *clubRepository) CreateClub(club *models.Tbl_club) error {
	return r.db.Create(club).Error
}

func (r *clubRepository) GetAllClubs() ([]models.Tbl_club, error) {
	var clubs []models.Tbl_club
	err := r.db.Find(&clubs).Error
	return clubs, err
}

func (r *clubRepository) GetClubById(id uint) (*models.Tbl_club, error) {
	var club models.Tbl_club
	err := r.db.First(&club, id).Error
	return &club, err
}

func (r *clubRepository) UpdateClub(club *models.Tbl_club) error {
	return r.db.Save(club).Error
}

func (r *clubRepository) DeleteClub(id uint) error {
	return r.db.Delete(&models.Tbl_club{}, id).Error
}
