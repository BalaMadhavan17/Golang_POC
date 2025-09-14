```go
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

	// Create the mdsListing table if it doesn't exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS mdsListing (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		comments TEXT,
		effectiveFrom DATE NOT NULL,
		effectiveTo DATE NOT NULL,
		isPPAgreed BOOLEAN DEFAULT FALSE,
		documentPath VARCHAR(255)
	);`)

	if err != nil {
		log.Fatal("❌ Failed to create mdsListing table: ", err)
	}

	log.Println("✅ Connected to MySQL Database")
}
```
