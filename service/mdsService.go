```
package service

import (
	"errors"
	"go-Beitler-api/repository"
	"time"
)

type MdsService interface {
	Delete(id int) error
	Create(name, comments string, effectiveFrom, effectiveTo time.Time, isPPAgreed bool) (int, error)
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

func (s *mdsService) Create(name, comments string, effectiveFrom, effectiveTo time.Time, isPPAgreed bool) (int, error) {
	if name == "" {
		return 0, errors.New("name is required")
	}
	
	if effectiveFrom.IsZero() {
		return 0, errors.New("effective from date is required")
	}
	
	if effectiveTo.IsZero() {
		return 0, errors.New("effective to date is required")
	}
	
	if effectiveTo.Before(effectiveFrom) {
		return 0, errors.New("effective to date cannot be earlier than effective from date")
	}
	
	entry := &repository.MdsEntry{
		Name:          name,
		Comments:      comments,
		EffectiveFrom: effectiveFrom,
		EffectiveTo:   effectiveTo,
		IsPPAgreed:    isPPAgreed,
		CreatedAt:     time.Now(),
	}
	
	return s.repo.Create(entry)
}

func (s *mdsService) GetAll() ([]repository.MdsEntry, error) {
	return s.repo.GetAll()
}
```
