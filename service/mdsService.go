```go
package service

import (
	"errors"
	"go-Beitler-api/repository"
	"time"
)

type MdsService interface {
	Delete(id int) error
	Create(entry repository.MdsEntry) (int, error)
	GetAll() ([]repository.MdsEntry, error)
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

func (s *mdsService) Create(entry repository.MdsEntry) (int, error) {
	// Validate required fields
	if entry.Name == "" {
		return 0, errors.New("name is required")
	}

	// Validate date logic
	if entry.EffectiveFrom.IsZero() {
		return 0, errors.New("effective from date is required")
	}

	if entry.EffectiveTo.IsZero() {
		return 0, errors.New("effective to date is required")
	}

	if entry.EffectiveTo.Before(entry.EffectiveFrom) {
		return 0, errors.New("effective to date must not be earlier than effective from date")
	}

	return s.repo.Create(entry)
}

func (s *mdsService) GetAll() ([]repository.MdsEntry, error) {
	return s.repo.GetAll()
}
```
