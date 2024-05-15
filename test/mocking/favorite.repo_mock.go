package mocking

import (
	"context"
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockFavoriteRepository struct {
	mock.Mock
}

func (m *MockFavoriteRepository) CreateFavorite(ctx context.Context, favorite *models.FavoriteRequest) (*config.Result, error) {
	ret := m.Called(ctx, favorite)

	var r0 *config.Result
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*config.Result)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockFavoriteRepository) DeleteFavorite(ctx context.Context, favoriteID string) error {
	ret := m.Called(ctx, favoriteID)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockFavoriteRepository) GetFavoritesByUserID(ctx context.Context, userID string) (*config.Result, error) {
	ret := m.Called(ctx, userID)

	var r0 *config.Result
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*config.Result)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
