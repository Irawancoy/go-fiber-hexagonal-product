package mocks

import (
	"go-fiber-hexagonal-product/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

// MockProductService adalah mock implementasi dari ProductService
type MockProductService struct {
	mock.Mock
}

// GetProduct adalah mock implementasi dari metode GetProduct
func (m *MockProductService) GetProduct(id string) (*domain.Product, error) {
	// Panggil metode yang di-mock dengan argumen id
	args := m.Called(id)
	// Jika hasil panggilan memiliki nilai, kembalikan nilai tersebut
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Product), args.Error(1)
	}
	// Jika hasil panggilan tidak memiliki nilai, kembalikan error
	return nil, args.Error(1)
}

// CreateProduct adalah mock implementasi dari metode CreateProduct
func (m *MockProductService) CreateProduct(product *domain.Product) error {
	// Panggil metode yang di-mock dengan argumen product
	args := m.Called(product)
	// Kembalikan error dari hasil panggilan
	return args.Error(0)
}

// UpdateProduct adalah mock implementasi dari metode UpdateProduct
func (m *MockProductService) UpdateProduct(product *domain.Product) error {
	// Panggil metode yang di-mock dengan argumen product
	args := m.Called(product)
	// Kembalikan error dari hasil panggilan
	return args.Error(0)
}

// DeleteProduct adalah mock implementasi dari metode DeleteProduct
func (m *MockProductService) DeleteProduct(id string) error {
	// Panggil metode yang di-mock dengan argumen id
	args := m.Called(id)
	// Kembalikan error dari hasil panggilan
	return args.Error(0)
}

// ListProducts adalah mock implementasi dari metode ListProducts
func (m *MockProductService) ListProducts() ([]*domain.Product, error) {
	// Panggil metode yang di-mock
	args := m.Called()
	// Jika hasil panggilan memiliki nilai, kembalikan nilai tersebut
	if args.Get(0) != nil {
		return args.Get(0).([]*domain.Product), args.Error(1)
	}
	// Jika hasil panggilan tidak memiliki nilai, kembalikan slice kosong untuk mencegah panic
	return []*domain.Product{}, args.Error(1)
}