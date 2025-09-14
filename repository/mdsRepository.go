package repository

import (
	"database/sql"
	"time"
)

type MdsEntry struct {
	Id             int
	Name           string
	Comments       string
	EffectiveFrom  time.Time
	EffectiveTo    time.Time
	IsPPAgreed     bool
	DocumentPath   string
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
	result, err := r.db.Exec("INSERT INTO mdsListing (name, comments, effective_from, effective_to, is_pp_agreed, document_path) VALUES (?, ?, ?, ?, ?, ?)",
		entry.Name, entry.Comments, entry.EffectiveFrom, entry.EffectiveTo, entry.IsPPAgreed, entry.DocumentPath)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *mdsRepository) GetAll() ([]MdsEntry, error) {
	rows, err := r.db.Query("SELECT id, name, comments, effective_from, effective_to, is_pp_agreed, document_path FROM mdsListing")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []MdsEntry
	for rows.Next() {
		var entry MdsEntry
		err := rows.Scan(&entry.Id, &entry.Name, &entry.Comments, &entry.EffectiveFrom, &entry.EffectiveTo, &entry.IsPPAgreed, &entry.DocumentPath)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
