package repository

import (
	"database/sql"
	"go-Beitler-api/models"
)

type MdsRepository interface {
	GetAll() ([]models.MdsListing, error)
	GetByID(id int) (*models.MdsListing, error)
	Create(mds models.MdsListing) error
	Update(mds models.MdsListing) error
	Delete(id int) error
}

type mdsRepository struct {
	db *sql.DB
}

func NewMdsRepository(db *sql.DB) MdsRepository {
	return &mdsRepository{db}
}

func (r *mdsRepository) GetAll() ([]models.MdsListing, error) {
	rows, err := r.db.Query("SELECT * FROM mdsListing")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mdsList []models.MdsListing
	for rows.Next() {
		var mds models.MdsListing
		err := rows.Scan(&mds.ID, &mds.MdsName, &mds.Comments, &mds.EffectiveFrom,
			&mds.EffectiveTo, &mds.IsPpAgreed, &mds.ReferenceNo, &mds.FilePath,
			&mds.CreatedAt, &mds.UpdatedAt)
		if err != nil {
			return nil, err
		}
		mdsList = append(mdsList, mds)
	}
	return mdsList, nil
}

func (r *mdsRepository) GetByID(id int) (*models.MdsListing, error) {
	var mds models.MdsListing
	err := r.db.QueryRow("SELECT * FROM mdsListing WHERE id=?", id).
		Scan(&mds.ID, &mds.MdsName, &mds.Comments, &mds.EffectiveFrom,
			&mds.EffectiveTo, &mds.IsPpAgreed, &mds.ReferenceNo, &mds.FilePath,
			&mds.CreatedAt, &mds.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &mds, nil
}

func (r *mdsRepository) Create(mds models.MdsListing) error {
	_, err := r.db.Exec(
		`INSERT INTO mdsListing (mdsName, comments, effectiveFrom, effectiveTo, isPpAgreed, referenceNo, filePath) 
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		mds.MdsName, mds.Comments, mds.EffectiveFrom, mds.EffectiveTo,
		mds.IsPpAgreed, mds.ReferenceNo, mds.FilePath,
	)
	return err
}

func (r *mdsRepository) Update(mds models.MdsListing) error {
	_, err := r.db.Exec(
		`UPDATE mdsListing SET mdsName=?, comments=?, effectiveFrom=?, effectiveTo=?, isPpAgreed=?, referenceNo=?, filePath=? 
		 WHERE id=?`,
		mds.MdsName, mds.Comments, mds.EffectiveFrom, mds.EffectiveTo,
		mds.IsPpAgreed, mds.ReferenceNo, mds.FilePath, mds.ID,
	)
	return err
}

func (r *mdsRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM mdsListing WHERE id=?", id)
	return err
}
