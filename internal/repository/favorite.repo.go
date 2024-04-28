package repository

import (
	"context"
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/models"
)

type FavoriteRepository interface {
	CreateFavorite(ctx context.Context, favorite *models.FavoriteRequest) (*config.Result, error)
	DeleteFavorite(ctx context.Context, favoriteID string) error
	GetFavoritesByUserID(ctx context.Context, userID string) (*config.Result, error)
}
