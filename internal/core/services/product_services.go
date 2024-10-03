package services

import (
	"go-fiber-hexagonal-product/internal/core/domain"
	"go-fiber-hexagonal-product/internal/core/ports"
)

type ProductService struct {
	mongoRepo ports.MongoProductRepository
	mysqlRepo  ports.MySQLProductRepository
}

func NewProductService(mongoRepo ports.MongoProductRepository, mysqlRepo ports.MySQLProductRepository) *ProductService {
	return &ProductService{
		mongoRepo: mongoRepo,
		mysqlRepo: mysqlRepo,
	}
}

func (s *ProductService) GetProduct(id string) (*domain.Product, error) {
	// Mengambil produk dari MongoDB
	return s.mongoRepo.GetProduct(id)
}

func (s *ProductService) CreateProduct(product *domain.Product) error {
	// Simpan ke MongoDB dan ambil ID yang dihasilkan
	productID, err := s.mongoRepo.CreateProduct(product)
	if err != nil {
		return err
	}

	// Update ID produk untuk MySQL
	product.ID = productID

	// Simpan ke MySQL
	if err := s.mysqlRepo.CreateProduct(product); err != nil {
		return err
	}

	return nil
}
func (s *ProductService) UpdateProduct(product *domain.Product) error {
	// Mengupdate produk di MongoDB
	if err := s.mongoRepo.UpdateProduct(product); err != nil {
		return err
	}
	
	// Mengupdate produk di MySQL
	return s.mysqlRepo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id string) error {
	// Hapus produk dari MongoDB
	if err := s.mongoRepo.DeleteProduct(id); err != nil {
		return err
	}

	// Hapus produk dari MySQL
	return s.mysqlRepo.DeleteProduct(id)
}

func (s *ProductService) ListProducts() ([]*domain.Product, error) {
	// Mendapatkan daftar produk dari MongoDB
	return s.mongoRepo.ListProducts()
}
