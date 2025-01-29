package handlers

import (
	"encoding/json"
	"motorcycle-sales/internal/domain/models"
	"motorcycle-sales/internal/domain/usecase"
	"net/http"
)

type OrderHandler struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase usecase.OrderUseCase) OrderHandler {
	return OrderHandler{orderUseCase: orderUseCase}
}

// CreateOrderHandler - создать новый заказ
func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.orderUseCase.CreateOrder(order)
	if err != nil {
		http.Error(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order created successfully"))
}
