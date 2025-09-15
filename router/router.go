package router

import (
	"go-Beitler-api/config"
	"go-Beitler-api/handler"
	"go-Beitler-api/repository"
	"go-Beitler-api/service"
	
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes() http.Handler {
	r := mux.NewRouter()
	
	// Get database connection
	db := config.GetDB()
	
	// Initialize repository
	mdsRepo := repository.NewMdsRepository(db)
	
	// Initialize service
	mdsService := service.NewMdsService(mdsRepo)
	
	// Initialize handler
	mdsHandler := handler.NewMdsHandler(mdsService)
	
	// API routes
	api := r.PathPrefix("/api").Subrouter()
	
	// MDS routes
	api.HandleFunc("/mds", mdsHandler.Create).Methods("POST")
	api.HandleFunc("/mds", mdsHandler.GetAll).Methods("GET")
	api.HandleFunc("/mds/{id:[0-9]+}", mdsHandler.Delete).Methods("DELETE")
	
	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	
	return c.Handler(r)
}
