package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"orderservice/service"
)

// GET order all
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	orders := service.GetAllOrders()
	json.NewEncoder(w).Encode(orders)
}

// GET order by user
func GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	userIDParam := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "invalid User ID", http.StatusBadRequest)
		return
	}

	orders := service.GetOrdersByUserID(userID)
	if len(orders) == 0 {
		http.Error(w, "Orders not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Contect-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
