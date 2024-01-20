// service/club_service.go
package service

import (
	"go-scoresheet/workflow/models"
	"go-scoresheet/workflow/repository"
)

type ClubService interface {
	CreateClub(club *models.Tbl_club) error
	GetAllClubs() ([]models.Tbl_club, error)
	GetClubById(id uint) (*models.Tbl_club, error)
	UpdateClub(club *models.Tbl_club) error
	DeleteClub(id uint) error
}

type clubService struct {
	clubRepo repository.ClubRepository
}

func NewClubService(clubRepo repository.ClubRepository) ClubService {
	return &clubService{
		clubRepo: clubRepo,
	}
}

func (s *clubService) CreateClub(club *models.Tbl_club) error {
	return s.clubRepo.CreateClub(club)
}

func (s *clubService) GetAllClubs() ([]models.Tbl_club, error) {
	return s.clubRepo.GetAllClubs()
}

func (s *clubService) GetClubById(id uint) (*models.Tbl_club, error) {
	return s.clubRepo.GetClubById(id)
}

func (s *clubService) UpdateClub(club *models.Tbl_club) error {
	return s.clubRepo.UpdateClub(club)
}

func (s *clubService) DeleteClub(id uint) error {
	return s.clubRepo.DeleteClub(id)
}
