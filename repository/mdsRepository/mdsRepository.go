package mdsRepository

import (
	"database/sql"
	"go-Beitler-api/model"
	"log"
	"time"
)

type MdsRepository interface {
	Create(mds *model.MdsEntry) (int, error)
	GetAll() ([]model.MdsEntry, error)
	Delete(id int) error
}

type mdsRepository struct {
	db *sql.DB
}

func NewMdsRepository(db *sql.DB) MdsRepository {
	return &mdsRepository{db}
}

func (r *mdsRepository) Create(mds *model.MdsEntry) (int, error) {
	mds.CreatedAt = time.Now()

	query := `INSERT INTO mds_entries (name, comments, effective_from, effective_to, is_pp_agreed, document_path, created_at)
			VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query, mds.Name, mds.Comments, mds.EffectiveFrom, mds.EffectiveTo,
		mds.IsPPAgreed, mds.DocumentPath, mds.CreatedAt)
	if err != nil {
		log.Printf("Error creating MDS entry: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *mdsRepository) GetAll() ([]model.MdsEntry, error) {
	query := `SELECT id, name, comments, effective_from, effective_to, is_pp_agreed, document_path, created_at
			FROM mds_entries ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error querying MDS entries: %v", err)
		return nil, err
	}
	defer rows.Close()

	var entries []model.MdsEntry
	for rows.Next() {
		var entry model.MdsEntry
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Comments, &entry.EffectiveFrom,
			&entry.EffectiveTo, &entry.IsPPAgreed, &entry.DocumentPath, &entry.CreatedAt)
		if err != nil {
			log.Printf("Error scanning MDS entry: %v", err)
			return nil, err
		}
		entries = append(entries, entry)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating MDS entries: %v", err)
		return nil, err
	}

	return entries, nil
}

func (r *mdsRepository) Delete(id int) error {
	query := `DELETE FROM mds_entries WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting MDS entry: %v", err)
		return err
	}
	return nil
}
