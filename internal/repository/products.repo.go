package repository

import (
	"context"
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.ProductRequest) (*config.Result, error)
	UpdateProduct(ctx context.Context, productID string, product *models.ProductRequest) error
	DeleteProduct(ctx context.Context, productID string) error
	GetAllProducts(ctx context.Context, foodType string, page, limit int) (*config.Result, error)
	GetProductByID(ctx context.Context, productID string) (*config.Result, error)
}
