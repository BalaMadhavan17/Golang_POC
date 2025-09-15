package main

import (
	"go-Beitler-api/config"
	"go-Beitler-api/router"
	"log"
	"net/http"
)

func main() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Setup router
	r := router.SetupRouter(db)

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
