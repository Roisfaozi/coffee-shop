package mocking

import (
	"context"
	"github.com/Roisfaozi/coffee-shop/config"
	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock type for ProductRepository
type MockProductRepository struct {
	mock.Mock
}

// CreateProduct is a mock for ProductRepository CreateProduct
func (m *MockProductRepository) CreateProduct(ctx context.Context, product *models.ProductRequest) (*config.Result, error) {
	ret := m.Called(ctx, product)

	// Mengambil nilai kembalian dari pemanggilan mock
	var r0 *config.Result
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*config.Result)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct is a mock for ProductRepository UpdateProduct
func (m *MockProductRepository) UpdateProduct(ctx context.Context, productID string, product *models.ProductRequest) error {
	ret := m.Called(ctx, productID, product)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

// DeleteProduct is a mock for ProductRepository DeleteProduct
func (m *MockProductRepository) DeleteProduct(ctx context.Context, productID string) error {
	ret := m.Called(ctx, productID)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

// GetAllProducts is a mock for ProductRepository GetAllProducts
func (m *MockProductRepository) GetAllProducts(ctx context.Context, foodType string, page, limit int) (*config.Result, error) {
	ret := m.Called(ctx, foodType, page, limit)

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

// GetProductByID is a mock for ProductRepository GetProductByID
func (m *MockProductRepository) GetProductByID(ctx context.Context, productID string) (*config.Result, error) {
	ret := m.Called(ctx, productID)
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

// Implement the other methods in a similar manner
