package mdsHandler

import (
	"encoding/json"
	"go-Beitler-api/model"
	"go-Beitler-api/service"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
)

type MdsHandler struct {
	service service.MdsService
}

func NewMdsHandler(service service.MdsService) *MdsHandler {
	return &MdsHandler{service}
}

func (h *MdsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var mds model.MdsEntry

	if err := json.NewDecoder(r.Body).Decode(&mds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := mds.Validate(); err != nil {
		response := map[string]string{"error": err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	id, err := h.service.Create(&mds)
	if err != nil {
		http.Error(w, "Failed to create MDS entry", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":      id,
		"message": "MDS entry has been saved successfully.",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *MdsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	entries, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve MDS entries", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}

func (h *MdsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete MDS entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
