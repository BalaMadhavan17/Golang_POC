package mdsHandler

import (
	"encoding/json"
	"go-Beitler-api/model"
	"go-Beitler-api/service"
	"net/http"
)

// MdsHandler struct
type MdsHandler struct {
	service service.MdsService
}

// NewMdsHandler creates a new MdsHandler
func NewMdsHandler(s service.MdsService) *MdsHandler {
	return &MdsHandler{service: s}
}

func (h *MdsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input service.MdsInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Map input to model
	entry := &model.MdsEntry{
		Name:          input.Name,
		Comments:      input.Comments,
		EffectiveFrom: input.EffectiveFrom,
		EffectiveTo:   input.EffectiveTo,
		IsPPAgreed:    input.IsPPAgreed,
		ReferenceNo:   input.ReferenceNo,
		DocumentPath:  input.DocumentPath,
	}

	// Call service
	if _, err := h.service.Create(entry); err != nil {
		http.Error(w, "Failed to create resource", http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"isSuccess": true,
		"message":   "Resource created successfully",
		"data":      entry,
	})
}

func (h *MdsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *MdsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete logic
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete endpoint"))
}
