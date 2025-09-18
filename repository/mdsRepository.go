package repository

import (
	"database/sql"
	"go-Beitler-api/model"
	"time"
)

type MdsRepository interface {
	Delete(id int) error
	Create(entry *model.MdsEntry) (int, error)
	// GetAll returns a paginated, sorted list of entries.
	// page is 1-based. pageSize controls number of items per page.
	// sortBy is the column name to sort on and sortOrder is "ASC" or "DESC".
	// returns (entries, totalItems, error)
	GetAll(page, pageSize int, sortBy, sortOrder string) ([]model.MdsEntry, int, error)
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
	result, err := r.db.Exec("INSERT INTO mdsListing (mdsName, comments, effectiveFrom, effectiveTo, isPpAgreed, filePath, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?)",
		entry.Name, entry.Comments, entry.EffectiveFrom, entry.EffectiveTo, entry.IsPPAgreed, entry.DocumentPath, entry.CreatedAt)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *mdsRepository) GetAll(page, pageSize int, sortBy, sortOrder string) ([]model.MdsEntry, int, error) {
	// Validate and map sortBy to actual column names to avoid SQL injection
	allowedSortColumns := map[string]string{
		"id":            "id",
		"name":          "mdsName",
		"effectiveFrom": "effectiveFrom",
		"effectiveTo":   "effectiveTo",
		"createdAt":     "createdAt",
		"updatedAt":     "updatedAt",
	}

	col, ok := allowedSortColumns[sortBy]
	if !ok {
		col = "id"
	}

	order := "ASC"
	if sortOrder == "DESC" || sortOrder == "desc" {
		order = "DESC"
	}

	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// Get total count first
	var total int
	if err := r.db.QueryRow("SELECT COUNT(*) FROM mdsListing").Scan(&total); err != nil {
		return nil, 0, err
	}

	query := "SELECT id, mdsName, comments, effectiveFrom, effectiveTo, isPpAgreed, filePath, createdAt, updatedAt FROM mdsListing ORDER BY " + col + " " + order + " LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var entries []model.MdsEntry
	for rows.Next() {
		var entry model.MdsEntry
		if err := rows.Scan(&entry.ID, &entry.Name, &entry.Comments, &entry.EffectiveFrom, &entry.EffectiveTo, &entry.IsPPAgreed, &entry.DocumentPath, &entry.CreatedAt, &entry.UpdatedAt); err != nil {
			return nil, 0, err
		}
		entries = append(entries, entry)
	}

	return entries, total, nil
}
