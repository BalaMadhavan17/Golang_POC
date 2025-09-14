```
package config

import (
	"database/sql"
	"fmt"
	"log"

	// "os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// Read from environment variables (export before running)
	dbUser := "balasubramanian_m"                                           //os.Getenv("DB_USER")
	dbPass := "4Vv03HRg818f"                                                //os.Getenv("DB_PASS")
	dbHost := "pocportal-dev.cvmt59aicyza.us-east-1.rds.amazonaws.com:3306" //os.Getenv("DB_HOST")
	dbName := "Beitler"                                                     //os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Error opening DB: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("❌ DB not reachable: ", err)
	}

	// Initialize required tables
	initTables()

	log.Println("✅ Connected to MySQL Database")
}

func initTables() {
	// Create mdsListing table if it doesn't exist
	query := `CREATE TABLE IF NOT EXISTS mdsListing (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		comments TEXT,
		effective_from DATE NOT NULL,
		effective_to DATE NOT NULL,
		is_pp_agreed BOOLEAN DEFAULT FALSE,
		document_path VARCHAR(500),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := DB.Exec(query)
	if err != nil {
		log.Printf("❌ Error creating table: %v", err)
	}
}
```
