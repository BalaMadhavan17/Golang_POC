```
package service

import (
	"errors"
	"go-Beitler-api/repository"
	"time"
)

type MdsService interface {
	Delete(id int) error
	Create(mds *repository.MDS) (int, error)
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

func (s *mdsService) Create(mds *repository.MDS) (int, error) {
	// Validate required fields
	if mds.Name == "" {
		return 0, errors.New("name is required")
	}

	if mds.EffectiveFrom.IsZero() {
		return 0, errors.New("effective from date is required")
	}

	if mds.EffectiveTo.IsZero() {
		return 0, errors.New("effective to date is required")
	}

	// Validate date logic
	if mds.EffectiveTo.Before(mds.EffectiveFrom) {
		return 0, errors.New("effective to date cannot be earlier than effective from date")
	}

	return s.repo.Create(mds)
}
```
