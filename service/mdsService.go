package service

import (
	"go-Beitler-api/model"
	"go-Beitler-api/repository"
	"time"
)

type MdsService interface {
	Create(mds *model.MdsEntry) (int, error)
	// GetAll returns paginated and sorted entries. page is 1-based.
	GetAll(page, pageSize int, sortBy, sortOrder string) ([]model.MdsEntry, error)
	Delete(id int) error
}

type mdsService struct {
	repo repository.MdsRepository
}

type MdsInput struct {
	Name          string    `json:"name"`
	Comments      string    `json:"comments"`
	EffectiveFrom time.Time `json:"effectiveFrom"`
	EffectiveTo   time.Time `json:"effectiveTo"`
	IsPPAgreed    bool      `json:"isPPAgreed"`
	DocumentPath  string    `json:"documentPath"`
	CreatedAt     time.Time `json:"createdAt"`
	ReferenceNo   string    `json:"referenceNo"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
}

func NewMdsService(repo repository.MdsRepository) MdsService {
	return &mdsService{repo}
}

func (s *mdsService) Create(mds *model.MdsEntry) (int, error) {
	return s.repo.Create(mds)
}
func (s *mdsService) GetAll(page, pageSize int, sortBy, sortOrder string) ([]model.MdsEntry, error) {
	return s.repo.GetAll(page, pageSize, sortBy, sortOrder)
}
func (s *mdsService) Delete(id int) error {
	return s.repo.Delete(id)
}
