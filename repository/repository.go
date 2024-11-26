package repository

import (
	"orderservice/model"
)

var orders = []model.Order{
	{ID: 1, UserID: 1, Total: "10"},
	{ID: 2, UserID: 2, Total: "40"},
	{ID: 3, UserID: 1, Total: "20"},
}

func GetAllOrders() []model.Order {
	return orders
}

func GetOrdersByUserID(userID int) []model.Order {
	var userOrders []model.Order
	for _, order := range orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders
}
