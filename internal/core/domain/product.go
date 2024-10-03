package domain

// Struktur data produk
type Product struct {
    // ID produk (unik)
    ID    string `json:"id" bson:"_id,omitempty" db:"product_id"`
    
    // Nama produk
    Name  string `json:"name" bson:"name" db:"product_name"`
    
    // Harga produk
    Price int    `json:"price" bson:"price" db:"price"`
    
    // Stok produk
    Stock int    `json:"stock" bson:"stock" db:"stock"`
}