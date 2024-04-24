package repository

import (
	"context"
	"time"

	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/jmoiron/sqlx"
)

type FavoriteRepositoryImpl struct {
	db *sqlx.DB
}

func (fr FavoriteRepositoryImpl) CreateFavorite(ctx context.Context, favorite *models.FavoriteRequest) (*models.FavoriteResponse, error) {
	tx, err := fr.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()
	query := `
        INSERT INTO favorite (product_id, user_id, created_at)
        VALUES ($1, $2, CURRENT_TIMESTAMP)
        RETURNING id, created_at
    `
	var favoriteID string
	var createdAt time.Time
	err = tx.QueryRowContext(ctx, query, favorite.ProductID, favorite.UserID).Scan(&favoriteID, &createdAt)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &models.FavoriteResponse{
		ID:        favoriteID,
		ProductID: favorite.ProductID,
		UserID:    favorite.UserID,
		CreatedAt: createdAt,
	}, nil
}

func (fr FavoriteRepositoryImpl) DeleteFavorite(ctx context.Context, favoriteID string) error {
	query := "DELETE FROM favorite WHERE id=$1"
	_, err := fr.db.ExecContext(ctx, query, favoriteID)
	if err != nil {
		return err
	}
	return nil
}

func (fr FavoriteRepositoryImpl) GetFavoritesByUserID(ctx context.Context, userID string) ([]*models.FavoriteResponse, error) {
	query := `
        SELECT id, product_id, user_id, created_at
        FROM favorite
        WHERE user_id = $1
    `

	rows, err := fr.db.QueryxContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []*models.FavoriteResponse
	for rows.Next() {
		var favorite models.FavoriteResponse
		err := rows.Scan(&favorite.ID, &favorite.ProductID, &favorite.UserID, &favorite.CreatedAt)
		if err != nil {
			return nil, err
		}
		favorites = append(favorites, &favorite)
	}

	return favorites, nil
}

func NewFavoriteRepositoryImpl(db *sqlx.DB) *FavoriteRepositoryImpl {
	return &FavoriteRepositoryImpl{db}
}
