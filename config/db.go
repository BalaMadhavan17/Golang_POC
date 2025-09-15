package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mds_db?parseTime=true")
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	// Create table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS mds_entries (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			comments TEXT,
			effective_from DATETIME NOT NULL,
			effective_to DATETIME NOT NULL,
			is_pp_agreed BOOLEAN DEFAULT FALSE,
			document_path VARCHAR(255),
			created_at DATETIME NOT NULL
		)
	`)
	if err != nil {
		log.Printf("Error creating table: %v", err)
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}
