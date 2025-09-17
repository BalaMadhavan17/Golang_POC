package main

import (
	"go-Beitler-api/config"
	"go-Beitler-api/handler"
	"go-Beitler-api/repository"
	"go-Beitler-api/router"
	"go-Beitler-api/service"
	"log"
	"net/http"
)

func main() {
	// Initialize database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create necessary tables if they don't exist
	setupDB(db)

	// Initialize repositories, services, and handlers
	mdsRepo := repository.NewMdsRepository(db)
	mdsService := service.NewMdsService(mdsRepo)
	mdsHandler := handler.NewMdsHandler(mdsService)

	// Set up routes
	r := router.SetupRoutes(mdsHandler)

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func setupDB(db *sql.DB) {
	// Create mdsListing table if it doesn't exist
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS mdsListing (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		mdsName TEXT NOT NULL,
		comments TEXT,
		effectiveFrom DATETIME NOT NULL,
		effectiveTo DATETIME NOT NULL,
		isPpAgreed BOOLEAN NOT NULL,
		filePath TEXT,
		referenceNo TEXT,
		createdAt DATETIME NOT NULL,
		updatedAt DATETIME
	)`,
	)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
