package models

import "time"

type Favorite struct {
	ID        string    `json:"id" db:"id"`
	ProductID string    `json:"product_id" db:"product_id"`
	UserID    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type FavoriteRequest struct {
	ProductID string `json:"product_id"`
	UserID    string `json:"user_id"`
}

type FavoriteResponse struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
