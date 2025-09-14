```
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
	DocumentPath  string
}

type MdsRepository interface {
	Delete(id int) error
	Create(entry *MdsEntry) (int, error)
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

func (r *mdsRepository) Create(entry *MdsEntry) (int, error) {
	result, err := r.db.Exec("INSERT INTO mdsListing (name, comments, effective_from, effective_to, is_pp_agreed, document_path) VALUES (?, ?, ?, ?, ?, ?)",
		entry.Name, entry.Comments, entry.EffectiveFrom, entry.EffectiveTo, entry.IsPPAgreed, entry.DocumentPath)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}
```
