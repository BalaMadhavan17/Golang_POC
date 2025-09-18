package model

import (
	"errors"
	"time"
)

type MdsEntry struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Comments      string    `json:"comments"`
	EffectiveFrom time.Time `json:"effectiveFrom"`
	EffectiveTo   time.Time `json:"effectiveTo"`
	IsPPAgreed    bool      `json:"isPPAgreed"`
	DocumentPath  string    `json:"documentPath"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
}

func (m *MdsEntry) Validate() error {
	if m.Name == "" {
		return errors.New("name is required")
	}
	if m.EffectiveFrom.IsZero() || m.EffectiveFrom.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return errors.New("effective from date is invalid or required")
	}
	if m.EffectiveTo.IsZero() || m.EffectiveTo.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		return errors.New("effective to date is invalid or required")
	}
	if m.EffectiveTo.Before(m.EffectiveFrom) {
		return errors.New("effective to date must not be earlier than effective from date")
	}
	return nil
}
