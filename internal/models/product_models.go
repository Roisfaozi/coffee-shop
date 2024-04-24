package models

import "time"

type Product struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Price       int       `json:"price" db:"price"`
	Currency    string    `json:"currency" db:"currency"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Sizes       []*Size   `json:"sizes,omitempty"`
}

type Size struct {
	ID       string `json:"id" db:"id"`
	SizeName string `json:"size_name" db:"size_name"`
}

type ProductSize struct {
	ProductID string `db:"product_id"`
	SizeID    string `db:"size_id"`
}

type ProductRequest struct {
	Name        string   `json:"name" binding:"required"`
	Price       int      `json:"price" binding:"required"`
	Currency    string   `json:"currency" binding:"required"`
	Description string   `json:"description"`
	ImageURL    string   `json:"image_url"`
	SizeIDs     []string `json:"size_ids"`
}

type ProductResponse struct {
	ID string `json:"id"`
}