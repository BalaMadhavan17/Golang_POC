package handler

import (
	"encoding/json"
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

func (h *MdsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Deleted successfully"})
}
