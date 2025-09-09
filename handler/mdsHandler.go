package handler

import (
	"encoding/json"
	"go-Beitler-api/models"
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

func (h *MdsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *MdsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	data, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *MdsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var mds models.MdsListing
	_ = json.NewDecoder(r.Body).Decode(&mds)
	err := h.service.Create(mds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Created successfully"})
}

func (h *MdsHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var mds models.MdsListing
	_ = json.NewDecoder(r.Body).Decode(&mds)
	mds.ID = id
	err := h.service.Update(mds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Updated successfully"})
}

func (h *MdsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted successfully"})
}
