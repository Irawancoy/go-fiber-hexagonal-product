package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-fiber-hexagonal-product/internal/adapters/handlers"
	"go-fiber-hexagonal-product/internal/core/domain"
	"go-fiber-hexagonal-product/internal/test/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestGetProduct adalah fungsi untuk menguji metode GetProduct
func TestGetProduct(t *testing.T) {
	// Buat mock product service
	mockProductService := new(mocks.MockProductService)

	// Buat fiber app
	app := fiber.New()

	// Buat product handler
	productHandler := handlers.NewProductHandler(mockProductService)

	// Tambahkan route untuk GetProduct
	app.Get("/products/:id", productHandler.GetProduct)

	// Buat mock product
	mockProduct := &domain.Product{
		ID:    "123",
		Name:  "Test Product",
		Price: 1000,
		Stock: 10,
	}

	// Test Success
	t.Run("Success", func(t *testing.T) {
		// Atur mock product service untuk mengembalikan mock product
		mockProductService.On("GetProduct", "123").Return(mockProduct, nil).Once()

		// Buat request untuk GetProduct
		req := httptest.NewRequest(http.MethodGet, "/products/123", nil)

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		// Decode response ke struct product
		var result domain.Product
		err = json.NewDecoder(resp.Body).Decode(&result)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa apakah result sama dengan mock product
		assert.Equal(t, mockProduct, &result)
	})

	// Test Not Found
	t.Run("Not Found", func(t *testing.T) {
		// Atur mock product service untuk mengembalikan error
		mockProductService.On("GetProduct", "456").Return(nil, errors.New("product not found")).Once()

		// Buat request untuk GetProduct
		req := httptest.NewRequest(http.MethodGet, "/products/456", nil)

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	// Periksa apakah mock product service telah dipanggil
	mockProductService.AssertExpectations(t)
}

// TestCreateProduct adalah fungsi untuk menguji metode CreateProduct
func TestCreateProduct(t *testing.T) {
	// Buat mock product service
	mockProductService := new(mocks.MockProductService)

	// Buat fiber app
	app := fiber.New()

	// Buat product handler
	productHandler := handlers.NewProductHandler(mockProductService)

	// Tambahkan route untuk CreateProduct
	app.Post("/product", productHandler.CreateProduct)

	// Buat mock product
	mockProduct := &domain.Product{
		Name:  "Test Product",
		Price: 1000,
		Stock: 10,
	}

	// Test Success
	t.Run("Success", func(t *testing.T) {
		// Atur mock product service untuk mengembalikan nil
		mockProductService.On("CreateProduct", mock.AnythingOfType("*domain.Product")).Return(nil).Once()

		// Buat request untuk CreateProduct
		body, _ := json.Marshal(mockProduct)
		req := httptest.NewRequest(http.MethodPost, "/product", bytes.NewReader(body))

		// Tambahkan header Content-Type
		req.Header.Set("Content-Type", "application/json")

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	})

	// Test Bad Request
	t.Run("Bad Request", func(t *testing.T) {
		// Buat request untuk CreateProduct dengan body yang tidak valid
		req := httptest.NewRequest(http.MethodPost, "/product", bytes.NewReader([]byte("invalid json")))

		// Tambahkan header Content-Type
		req.Header.Set("Content-Type", "application/json")

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	// Periksa apakah mock product service telah dipanggil
	mockProductService.AssertExpectations(t)
}

// TestUpdateProduct adalah fungsi untuk menguji metode UpdateProduct
func TestUpdateProduct(t *testing.T) {
	// Buat mock product service
	mockProductService := new(mocks.MockProductService)

	// Buat fiber app
	app := fiber.New()

	// Buat product handler
	productHandler := handlers.NewProductHandler(mockProductService)

	// Tambahkan route untuk UpdateProduct
	app.Put("/product/:id", productHandler.UpdateProduct)

	// Buat mock product
	mockProduct := &domain.Product{
		ID:    "123",
		Name:  "Updated Product",
		Price: 1500,
		Stock: 15,
	}

	// Test Success
	t.Run("Success", func(t *testing.T) {
		// Atur mock product service untuk mengembalikan nil
		mockProductService.On("UpdateProduct", mock.AnythingOfType("*domain.Product")).Return(nil).Once()

		// Buat request untuk UpdateProduct
		body, _ := json.Marshal(mockProduct)
		req := httptest.NewRequest(http.MethodPut, "/product/123", bytes.NewReader(body))

		// Tambahkan header Content-Type
		req.Header.Set("Content-Type", "application/json")

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	// Test Not Found
	t.Run("Not Found", func(t *testing.T) {
		// Atur mock product service untuk mengembalikan error
		mockProductService.On("UpdateProduct", mock.AnythingOfType("*domain.Product")).Return(errors.New("product not found")).Once()

		// Buat request untuk UpdateProduct
		body, _ := json.Marshal(mockProduct)
		req := httptest.NewRequest(http.MethodPut, "/product/456", bytes.NewReader(body))

		// Tambahkan header Content-Type
		req.Header.Set("Content-Type", "application/json")

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	// Periksa apakah mock product service telah dipanggil
	mockProductService.AssertExpectations(t)
}

// TestDeleteProduct adalah fungsi untuk menguji metode DeleteProduct
func TestDeleteProduct(t *testing.T) {
	// Buat mock product service
	mockProductService := new(mocks.MockProductService)

	// Buat fiber app
	app := fiber.New()

	// Buat product handler
	productHandler := handlers.NewProductHandler(mockProductService)

	// Tambahkan route untuk DeleteProduct
	app.Delete("/product/:id", productHandler.DeleteProduct)

	// Test Success
	t.Run("Success", func(t *testing.T) {
		// Atur mock product service untuk mengembalikan nil
		mockProductService.On("DeleteProduct", "123").Return(nil).Once()

		// Buat request untuk DeleteProduct
		req := httptest.NewRequest(http.MethodDelete, "/product/123", nil)

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	// Test Not Found
	t.Run("Not Found", func(t *testing.T) {
		// Atur mock product service untuk mengembalikan error
		mockProductService.On("DeleteProduct", "456").Return(errors.New("product not found")).Once()

		// Buat request untuk DeleteProduct
		req := httptest.NewRequest(http.MethodDelete, "/product/456", nil)

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	// Periksa apakah mock product service telah dipanggil
	mockProductService.AssertExpectations(t)
}

// TestListProducts adalah fungsi untuk menguji metode ListProducts
func TestListProducts(t *testing.T) {
 // Buat mock product service
	mockProductService := new(mocks.MockProductService)

	// Buat fiber app
	app := fiber.New()

	// Buat product handler
	productHandler := handlers.NewProductHandler(mockProductService)

	// Tambahkan route untuk ListProducts
	app.Get("/product", productHandler.ListProducts)

	// Buat mock products
	mockProducts := []*domain.Product{
		{
			ID:    "123",
			Name:  "Test Product",
			Price: 1000,
			Stock: 10,
		},
		{
			ID:    "456",
			Name:  "Test Product 2",
			Price: 2000,
			Stock: 20,
		},
	}

	// Test Success
	t.Run("Success", func(t *testing.T) {
		// Atur mock product service untuk mengembalikan mock products
		mockProductService.On("ListProducts").Return(mockProducts, nil).Once()

		// Buat request untuk ListProducts
		req := httptest.NewRequest(http.MethodGet, "/product", nil)

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		// Decode response ke struct products
		var result []*domain.Product
		err = json.NewDecoder(resp.Body).Decode(&result)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa apakah result sama dengan mock products
		assert.Equal(t, mockProducts, result)
	})

	// Test Not Found
	t.Run("Not Found", func(t *testing.T) {
		// Atur mock product service untuk mengembalikan error
		mockProductService.On("ListProducts").Return(nil, errors.New("products not found")).Once()

		// Buat request untuk ListProducts
		req := httptest.NewRequest(http.MethodGet, "/product", nil)

		// Kirim request ke fiber app
		resp, err := app.Test(req)

		// Periksa apakah ada error
		assert.NoError(t, err)

		// Periksa status code
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	})

	// Periksa apakah mock product service telah dipanggil
	mockProductService.AssertExpectations(t)
}