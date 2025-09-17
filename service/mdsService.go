```go
package service

import (
	"go-Beitler-api/model"
	"go-Beitler-api/repository"
	"time"
	"bytes"
	"encoding/csv"
)

type MdsService interface {
	Create(mds *model.MdsEntry) (int, error)
	GetAll() ([]model.MdsEntry, error)
	Delete(id int) error
	GetTemplate() ([]byte, error)
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

func (s *mdsService) GetAll() ([]model.MdsEntry, error) {
	return s.repo.GetAll()
}

func (s *mdsService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *mdsService) GetTemplate() ([]byte, error) {
	// Create a CSV template for MDS entries
	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)

	// Write header
	header := []string{"MDS Name", "Comments", "Effective From (YYYY-MM-DD)", "Effective To (YYYY-MM-DD)", "Is PP Agreed (true/false)", "Reference Number"}
	if err := writer.Write(header); err != nil {
		return nil, err
	}

	// Write example row
	exampleRow := []string{"Example MDS", "Example comment", time.Now().Format("2006-01-02"), time.Now().AddDate(1, 0, 0).Format("2006-01-02"), "true", "REF-12345"}
	if err := writer.Write(exampleRow); err != nil {
		return nil, err
	}

	writer.Flush()
	return buffer.Bytes(), nil
}
```
