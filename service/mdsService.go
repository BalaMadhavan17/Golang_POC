package service

import (
	"go-Beitler-api/models"
	"go-Beitler-api/repository"
)

type MdsService interface {
	GetAll() ([]models.MdsListing, error)
	GetByID(id int) (*models.MdsListing, error)
	Create(mds models.MdsListing) error
	Update(mds models.MdsListing) error
	Delete(id int) error
}

type mdsService struct {
	repo repository.MdsRepository
}

func NewMdsService(repo repository.MdsRepository) MdsService {
	return &mdsService{repo}
}

func (s *mdsService) GetAll() ([]models.MdsListing, error) {
	return s.repo.GetAll()
}

func (s *mdsService) GetByID(id int) (*models.MdsListing, error) {
	return s.repo.GetByID(id)
}

func (s *mdsService) Create(mds models.MdsListing) error {
	return s.repo.Create(mds)
}

func (s *mdsService) Update(mds models.MdsListing) error {
	return s.repo.Update(mds)
}

func (s *mdsService) Delete(id int) error {
	return s.repo.Delete(id)
}
