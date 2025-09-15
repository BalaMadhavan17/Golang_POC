package config

import (
	"database/sql"
	"log"
	"sync"
	
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var once sync.Once

func GetDB() *sql.DB {
	once.Do(func() {
		conn, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/mds_db?parseTime=true")
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		
		err = conn.Ping()
		if err != nil {
			log.Fatalf("Could not establish database connection: %v", err)
		}
		
		// Create table if not exists
		_, err = conn.Exec(`
			CREATE TABLE IF NOT EXISTS mds_entries (
				id INT AUTO_INCREMENT PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				comments TEXT,
				effective_from DATETIME NOT NULL,
				effective_to DATETIME NOT NULL,
				is_pp_agreed BOOLEAN DEFAULT FALSE,
				document_path VARCHAR(255),
				created_at DATETIME
			)
		`)
		if err != nil {
			log.Fatalf("Could not create table: %v", err)
		}
		
		db = conn
	})
	
	return db
}
