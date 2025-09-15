package router

import (
	"go-Beitler-api/handler"
	"go-Beitler-api/service"

	"github.com/gorilla/mux"
)

func SetupRouter(mdsHandler *handler.MdsHandler) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	mdsRoutes := api.PathPrefix("/mds").Subrouter()

	mdsRoutes.HandleFunc("", mdsHandler.Create).Methods("POST")
	mdsRoutes.HandleFunc("", mdsHandler.GetAll).Methods("GET")
	mdsRoutes.HandleFunc("/{id:[0-9]+}", mdsHandler.Delete).Methods("DELETE")

	return r
}
