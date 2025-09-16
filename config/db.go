package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var once sync.Once

func GetDB() *sql.DB {
	dbUser := "balasubramanian_m"                                           //os.Getenv("DB_USER")
	dbPass := "4Vv03HRg818f"                                                //os.Getenv("DB_PASS")
	dbHost := "pocportal-dev.cvmt59aicyza.us-east-1.rds.amazonaws.com:3306" //os.Getenv("DB_HOST")
	dbName := "Beitler"                                                     //os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=UTC", dbUser, dbPass, dbHost, dbName)
	once.Do(func() {
		conn, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		err = conn.Ping()
		if err != nil {
			log.Fatalf("Could not establish database connection: %v", err)
		}

		// Create table if not exists
		// _, err = conn.Exec(`
		// 	CREATE TABLE IF NOT EXISTS mds_entries (
		// 		id INT AUTO_INCREMENT PRIMARY KEY,
		// 		name VARCHAR(255) NOT NULL,
		// 		comments TEXT,
		// 		effective_from DATETIME NOT NULL,
		// 		effective_to DATETIME NOT NULL,
		// 		is_pp_agreed BOOLEAN DEFAULT FALSE,
		// 		document_path VARCHAR(255),
		// 		created_at DATETIME
		// 	)
		// `)
		// if err != nil {
		// 	log.Fatalf("Could not create table: %v", err)
		// }

		db = conn
	})

	return db
}
