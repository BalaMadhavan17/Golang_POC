package mdsHandler

import (
	"encoding/json"
	"go-Beitler-api/model"
	"go-Beitler-api/service"
	"net/http"
	"strconv"
	"strings"
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
	// Parse query params
	q := r.URL.Query()
	page := 1
	pageSize := 10
	sortBy := "name"
	sortOrder := "ASC"

	if p := q.Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if ps := q.Get("pageSize"); ps != "" {
		if v, err := strconv.Atoi(ps); err == nil && v > 0 {
			pageSize = v
		}
	}
	if sb := q.Get("sortBy"); sb != "" {
		sortBy = strings.TrimSpace(sb)
	}
	if so := q.Get("sortOrder"); so != "" {
		so = strings.ToUpper(strings.TrimSpace(so))
		if so == "ASC" || so == "DESC" {
			sortOrder = so
		}
	}

	data, total, err := h.service.GetAll(page, pageSize, sortBy, sortOrder)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	totalPages := 0
	if pageSize > 0 {
		totalPages = (total + pageSize - 1) / pageSize
	}

	envelope := map[string]interface{}{
		"data": data,
		"meta": map[string]interface{}{
			"totalItems": total,
			"totalPages": totalPages,
			"page":       page,
			"pageSize":   pageSize,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(envelope); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *MdsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete logic
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete endpoint"))
}
