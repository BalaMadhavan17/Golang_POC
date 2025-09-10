package service

import (
	"go-Beitler-api/repository"
)

type MdsService interface {
	Delete(id int) error
}

type mdsService struct {
	repo repository.MdsRepository
}

func NewMdsService(repo repository.MdsRepository) MdsService {
	return &mdsService{repo}
}

func (s *mdsService) Delete(id int) error {
	return s.repo.Delete(id)
}
