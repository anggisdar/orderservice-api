package model

// Order struct

type Order struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Total  string `json:"total"`
}
