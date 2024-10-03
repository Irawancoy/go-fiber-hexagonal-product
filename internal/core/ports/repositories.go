package ports

import "go-fiber-hexagonal-product/internal/core/domain"

// Interface untuk repository produk MongoDB
type MongoProductRepository interface {
    // Mendapatkan produk berdasarkan ID
    GetProduct(id string) (*domain.Product, error)
    
    // Membuat produk baru dan mengembalikan ID dari MongoDB
    CreateProduct(product *domain.Product) (string, error)
    
    // Mengupdate produk yang sudah ada
    UpdateProduct(product *domain.Product) error
    
    // Menghapus produk berdasarkan ID
    DeleteProduct(id string) error
    
    // Mendapatkan daftar produk
    ListProducts() ([]*domain.Product, error)
}

// Interface untuk repository produk MySQL
type MySQLProductRepository interface {
    // Mendapatkan produk berdasarkan ID
    GetProduct(id string) (*domain.Product, error)
    
    // Membuat produk baru
    CreateProduct(product *domain.Product) error
    
    // Mengupdate produk yang sudah ada
    UpdateProduct(product *domain.Product) error
    
    // Menghapus produk berdasarkan ID
    DeleteProduct(id string) error
    
    // Mendapatkan daftar produk
    ListProducts() ([]*domain.Product, error)
}