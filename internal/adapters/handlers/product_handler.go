package handlers

import (
	"go-fiber-hexagonal-product/internal/core/domain"
	"go-fiber-hexagonal-product/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

// Handler untuk produk
type ProductHandler struct {
	productService ports.ProductService
}

// Membuat instance baru dari ProductHandler
func NewProductHandler(productService ports.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// Mendapatkan produk berdasarkan ID
func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.productService.GetProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(product)
}

// Membuat produk baru
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	product := new(domain.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.productService.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Return product with status 201 Created
	return c.Status(fiber.StatusCreated).JSON(product)
}

// Mengupdate produk yang sudah ada
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(domain.Product)

	// Parse body dari request ke struct product
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Set ID produk dari parameter URL
	product.ID = id

	// Pastikan kita tidak mengubah field _id saat update
	if err := h.productService.UpdateProduct(product); err != nil {
		// Periksa apakah error disebabkan karena produk tidak ditemukan
		if err.Error() == "product not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Kembalikan response dengan data produk yang diupdate
	return c.JSON(product)
}

// Menghapus produk berdasarkan ID
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.productService.DeleteProduct(id); err != nil {
		// Periksa apakah error disebabkan karena produk tidak ditemukan
		if err.Error() == "product not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}

// Mendapatkan daftar produk
func (h *ProductHandler) ListProducts(c *fiber.Ctx) error {
	products, err := h.productService.ListProducts()
	if err != nil {
		if err.Error() == "products not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Products not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Jika products adalah nil, kembalikan slice kosong agar tidak terjadi panic
	if products == nil {
		products = []*domain.Product{}
	}

	return c.JSON(products)
}