package repository

import (
	"context"
	"github.com/Roisfaozi/coffee-shop/internal/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.ProductRequest) (*models.ProductResponse, error)
	UpdateProduct(ctx context.Context, productID string, product *models.ProductRequest) error
	DeleteProduct(ctx context.Context, productID string) error
	GetAllProducts(ctx context.Context, foodType string, page, limit int) ([]*models.Product, error)
	GetProductByID(ctx context.Context, productID string) (*models.Product, error)
}
