```go
package service

import (
	"errors"
	"go-Beitler-api/repository"
	"time"
)

type MdsService interface {
	Delete(id int) error
	Create(name, comments string, effectiveFrom, effectiveTo time.Time, isPPAgreed bool, documentPath string) (int, error)
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

func (s *mdsService) Create(name, comments string, effectiveFrom, effectiveTo time.Time, isPPAgreed bool, documentPath string) (int, error) {
	// Validate required fields
	if name == "" {
		return 0, errors.New("name is required")
	}

	// Validate date logic
	if !effectiveFrom.IsZero() && !effectiveTo.IsZero() && effectiveTo.Before(effectiveFrom) {
		return 0, errors.New("effective to date must not be earlier than effective from date")
	}

	// Create entry
	entry := repository.MdsEntry{
		Name:          name,
		Comments:      comments,
		EffectiveFrom: effectiveFrom,
		EffectiveTo:   effectiveTo,
		IsPPAgreed:    isPPAgreed,
		DocumentPath:  documentPath,
	}

	return s.repo.Create(entry)
}

func (s *mdsService) GetAll() ([]repository.MdsEntry, error) {
	return s.repo.GetAll()
}
```
