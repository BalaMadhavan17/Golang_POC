package router

import (
	"go-Beitler-api/handler"
	"github.com/gorilla/mux"
)

func SetupRoutes(mdsHandler *handler.MdsHandler) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/mds", mdsHandler.Create).Methods("POST")
	api.HandleFunc("/mds", mdsHandler.GetAll).Methods("GET")
	api.HandleFunc("/mds/{id}", mdsHandler.Delete).Methods("DELETE")
	api.HandleFunc("/mds/template", mdsHandler.GetTemplate).Methods("GET")

	return r
}
