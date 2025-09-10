package repository

import (
	"database/sql"
)

type MdsRepository interface {
	Delete(id int) error
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
