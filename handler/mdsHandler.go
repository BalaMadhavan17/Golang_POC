package handler

import (
	"encoding/json"
	"go-Beitler-api/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type CreateMdsRequest struct {
	Name          string    `json:"name"`
	Comments      string    `json:"comments"`
	EffectiveFrom string    `json:"effectiveFrom"`
	EffectiveTo   string    `json:"effectiveTo"`
	IsPPAgreed    bool      `json:"isPPAgreed"`
}

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

func (h *MdsHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Could not parse form", http.StatusBadRequest)
		return
	}

	// Get form values
	var req CreateMdsRequest
	req.Name = r.FormValue("name")
	req.Comments = r.FormValue("comments")
	req.EffectiveFrom = r.FormValue("effectiveFrom")
	req.EffectiveTo = r.FormValue("effectiveTo")
	isPPAgreedStr := r.FormValue("isPPAgreed")
	req.IsPPAgreed = isPPAgreedStr == "true"

	// Validation for required fields
	if req.Name == "" || req.EffectiveFrom == "" || req.EffectiveTo == "" {
		http.Error(w, "Please fill in all required fields", http.StatusBadRequest)
		return
	}

	// Parse dates
	effectiveFrom, err := time.Parse("2006-01-02", req.EffectiveFrom)
	if err != nil {
		http.Error(w, "Invalid effective from date format", http.StatusBadRequest)
		return
	}

	effectiveTo, err := time.Parse("2006-01-02", req.EffectiveTo)
	if err != nil {
		http.Error(w, "Invalid effective to date format", http.StatusBadRequest)
		return
	}

	// Handle file upload
	var documentPath string
	file, handler, err := r.FormFile("document")
	if err == nil {
		defer file.Close()

		// Create uploads directory if it doesn't exist
		uploadDir := "./uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, 0755)
		}

		// Create a unique filename
		documentPath = filepath.Join(uploadDir, handler.Filename)

		// Save the file
		dst, err := os.Create(documentPath)
		if err != nil {
			http.Error(w, "Failed to save the uploaded file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Failed to save the uploaded file", http.StatusInternalServerError)
			return
		}
	}

	// Save to database through service
	id, err := h.service.Create(req.Name, req.Comments, effectiveFrom, effectiveTo, req.IsPPAgreed, documentPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "MDS entry has been saved successfully.",
		"id":      id,
	})
}

func (h *MdsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	entries, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
