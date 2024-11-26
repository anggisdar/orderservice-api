package service

import (
	"orderservice/model"
	"orderservice/repository"
)

func GetAllOrders() []model.Order {
	return repository.GetAllOrders()
}

func GetOrdersByUserID(userID int) []model.Order {
	return repository.GetOrdersByUserID(userID)
}
