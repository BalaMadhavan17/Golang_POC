package repository

import (
	"database/sql"
	"time"
)

type MdsEntry struct {
	ID            int
	Name          string
	Comments      string
	EffectiveFrom time.Time
	EffectiveTo   time.Time
	IsPPAgreed    bool
}

type MdsRepository interface {
	Delete(id int) error
	Create(entry MdsEntry) (int, error)
	GetAll() ([]MdsEntry, error)
}

type mdsRepository struct {
	db *sql.DB
}

func NewMdsRepository(db *sql.DB) MdsRepository {
	return &mdsRepository{db}
}

func (r *mdsRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM mdsListing WHERE id=?", id)
	return err
}

func (r *mdsRepository) Create(entry MdsEntry) (int, error) {
	result, err := r.db.Exec("INSERT INTO mdsListing (name, comments, effective_from, effective_to, is_pp_agreed) VALUES (?, ?, ?, ?, ?)",
		entry.Name, entry.Comments, entry.EffectiveFrom, entry.EffectiveTo, entry.IsPPAgreed)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *mdsRepository) GetAll() ([]MdsEntry, error) {
	rows, err := r.db.Query("SELECT id, name, comments, effective_from, effective_to, is_pp_agreed FROM mdsListing")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := []MdsEntry{}
	for rows.Next() {
		entry := MdsEntry{}
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Comments, &entry.EffectiveFrom, &entry.EffectiveTo, &entry.IsPPAgreed)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
