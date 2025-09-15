package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./mds.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create tables if they don't exist
	createTablesQuery := `
	CREATE TABLE IF NOT EXISTS mds_entries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		comments TEXT,
		effective_from DATETIME NOT NULL,
		effective_to DATETIME NOT NULL,
		is_pp_agreed BOOLEAN DEFAULT 0,
		document_path TEXT,
		created_at DATETIME
	);
	`
	_, err = db.Exec(createTablesQuery)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	return db
}
