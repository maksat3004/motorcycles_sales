package handlers

import (
	"encoding/json"
	"motorcycle-sales/internal/domain/models"
	"motorcycle-sales/internal/domain/usecase"
	"net/http"
)

type MotorcycleHandler struct {
	motorcycleUseCase usecase.MotorcycleUseCase
}

func NewMotorcycleHandler(motorcycleUseCase usecase.MotorcycleUseCase) MotorcycleHandler {
	return MotorcycleHandler{motorcycleUseCase: motorcycleUseCase}
}

// GetAllMotorcyclesHandler - получить все мотоциклы
func (h *MotorcycleHandler) GetAllMotorcyclesHandler(w http.ResponseWriter, r *http.Request) {
	motorcycles, err := h.motorcycleUseCase.GetAllMotorcycles()
	if err != nil {
		http.Error(w, "Failed to fetch motorcycles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(motorcycles)
}

// AddMotorcycleHandler - добавить новый мотоцикл
func (h *MotorcycleHandler) AddMotorcycleHandler(w http.ResponseWriter, r *http.Request) {
	var motorcycle models.Motorcycle
	err := json.NewDecoder(r.Body).Decode(&motorcycle)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.motorcycleUseCase.AddMotorcycle(motorcycle)
	if err != nil {
		http.Error(w, "Failed to add motorcycle: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Motorcycle added successfully"))
}
