package main

import (
	"log"
	"net/http"

	"orderservice/handler"
)

func main() {
	http.HandleFunc("/orders", handler.GetOrders)
	http.HandleFunc("/orders/user", handler.GetOrdersByUserID)

	log.Println("Starting server on :8002")
	if err := http.ListenAndServe(":8002", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
