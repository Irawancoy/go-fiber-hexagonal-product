package repositories

import (
	"database/sql"
	"go-fiber-hexagonal-product/internal/core/domain"
	"log"
)

// Repository produk MySQL
type MysqlProductRepository struct {
	db *sql.DB
}

// Membuat instance baru dari MysqlProductRepository
func NewMySQLProductRepository(db *sql.DB) *MysqlProductRepository {
	return &MysqlProductRepository{db: db}
}

// Mendapatkan produk berdasarkan ID
func (r *MysqlProductRepository) GetProduct(id string) (*domain.Product, error) {
	var product domain.Product
	err := r.db.QueryRow("SELECT product_id, product_name, price, stock FROM product WHERE product_id = ?", id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			// Jika produk tidak ditemukan, return nil dan error nil
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

// Membuat produk baru
func (r *MysqlProductRepository) CreateProduct(product *domain.Product) error {
	_, err := r.db.Exec("INSERT INTO product (product_id, product_name, price, stock) VALUES (?, ?, ?, ?)", product.ID, product.Name, product.Price, product.Stock)
	if err != nil {
		log.Printf("Gagal membuat produk di MySQL: %v", err)
		return err
	}
	return nil
}

// Mengupdate produk yang sudah ada
func (r *MysqlProductRepository) UpdateProduct(product *domain.Product) error {
	// Cek apakah produk ada di MySQL
	_, err := r.GetProduct(product.ID)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("UPDATE product SET product_name = ?, price = ?, stock = ? WHERE product_id = ?", product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		log.Printf("Gagal mengupdate produk di MySQL: %v", err)
		return err
	}
	return nil
}

// Menghapus produk berdasarkan ID
func (r *MysqlProductRepository) DeleteProduct(id string) error {
	_, err := r.db.Exec("DELETE FROM product WHERE product_id = ?", id)
	if err != nil {
		log.Printf("Gagal menghapus produk: %v", err)
		return err
	}
	return nil
}

// Mendapatkan daftar produk
func (r *MysqlProductRepository) ListProducts() ([]*domain.Product, error) {
	rows, err := r.db.Query("SELECT product_id, product_name, price, stock FROM product")
	if err != nil {
		log.Printf("Gagal mendapatkan daftar produk: %v", err)
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			log.Printf("Gagal scan produk: %v", err)
			return nil, err
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return products, nil
}