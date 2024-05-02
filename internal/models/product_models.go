package models

import "time"

type Product struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Price       int       `json:"price" db:"price"`
	Currency    string    `json:"currency" db:"currency"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	Category    string    `json:"category" db:"category"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Sizes       []*Size   `json:"sizes,omitempty" db:"-"`
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
	Name        string   `json:"name" form:"name" db:"name" binding:"required" valid:"type(string),required"`
	Price       int      `json:"price" form:"price" db:"price" binding:"required" valid:"type(int),required"`
	Currency    string   `json:"currency" form:"currency" db:"currency" binding:"required" valid:"type(string),required"`
	Description string   `json:"description" form:"description" db:"description" valid:"type(string)"`
	ImageURL    string   `json:"image_url" form:"image_url" db:"image_url"  valid:"type(string)"`
	Category    string   `json:"category" form:"category" db:"category" valid:"type(string)"`
	SizeIDs     []string `json:"size_ids"`
}

type ProductResponse struct {
	ID string `json:"id"`
}
