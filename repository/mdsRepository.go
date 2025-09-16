package repository

import (
	"database/sql"
	"go-Beitler-api/model"
	"time"
)

type MdsRepository interface {
	Delete(id int) error
	Create(entry *model.MdsEntry) (int, error)
	GetAll() ([]model.MdsEntry, error)
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

func (r *mdsRepository) Create(entry *model.MdsEntry) (int, error) {
	entry.CreatedAt = time.Now()
	result, err := r.db.Exec("INSERT INTO mdsListing (mdsName, comments, effectiveFrom, effectiveTo, isPpAgreed, filePath, referenceNo, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		entry.Name, entry.Comments, entry.EffectiveFrom, entry.EffectiveTo, entry.IsPPAgreed, entry.DocumentPath, entry.ReferenceNo, entry.CreatedAt)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *mdsRepository) GetAll() ([]model.MdsEntry, error) {
	rows, err := r.db.Query("SELECT id, mdsName, comments, effectiveFrom, effectiveTo, isPpAgreed, referenceNo, filePath, createdAt, updatedAt FROM mdsListing")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []model.MdsEntry
	for rows.Next() {
		var entry model.MdsEntry
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Comments, &entry.EffectiveFrom, &entry.EffectiveTo, &entry.IsPPAgreed, &entry.ReferenceNo, &entry.DocumentPath, &entry.CreatedAt, &entry.UpdatedAt)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
