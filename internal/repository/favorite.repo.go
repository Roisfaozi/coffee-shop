package repository

import (
	"context"
	"github.com/Roisfaozi/coffee-shop/internal/models"
)

type FavoriteRepository interface {
	CreateFavorite(ctx context.Context, favorite *models.FavoriteRequest) (*models.FavoriteResponse, error)
	DeleteFavorite(ctx context.Context, favoriteID string) error
	GetFavoritesByUserID(ctx context.Context, userID string) ([]*models.FavoriteResponse, error)
}
