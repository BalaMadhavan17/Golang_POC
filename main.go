package main

import (
	"fmt"
	"log"
	"net/http"

	"go-Beitler-api/config"
	"go-Beitler-api/handler"
	"go-Beitler-api/repository"
	"go-Beitler-api/router"
	"go-Beitler-api/service"
)

func main() {
	// Initialize database connection
	db := config.InitDB()
	defer db.Close()

	// Initialize repository, service, and handler
	mdsRepo := repository.NewMdsRepository(db)
	mdsService := service.NewMdsService(mdsRepo)
	mdsHandler := handler.NewMdsHandler(mdsService)

	// Setup router
	r := router.SetupRouter(mdsHandler)

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
