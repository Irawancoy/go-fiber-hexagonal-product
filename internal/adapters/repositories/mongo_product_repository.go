package repositories

import (
	"context"
	"go-fiber-hexagonal-product/internal/core/domain"
	"log"
	"time" // Package ini digunakan untuk mengukur durasi/kecepatan koneksi ke MongoDB

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository produk MongoDB
type MongoProductRepository struct {
	collection *mongo.Collection
}

// Membuat instance baru dari MongoProductRepository
func NewMongoProductRepository(collection *mongo.Collection) *MongoProductRepository {
	return &MongoProductRepository{
		collection: collection,
	}
}

// Mendapatkan produk berdasarkan ID
func (r *MongoProductRepository) GetProduct(id string) (*domain.Product, error) {
	start := time.Now() // Mulai pengukuran waktu
	var product domain.Product
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// Mengambil produk dari MongoDB berdasarkan ID
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	// Menghitung durasi waktu yang dihabiskan untuk query
	log.Printf("GetProduct duration: %v", time.Since(start)) // Waktu yang diambil dari awal hingga selesai query
	return &product, nil
}

// Membuat produk baru
func (r *MongoProductRepository) CreateProduct(product *domain.Product) (string, error) {
	start := time.Now() // Mulai pengukuran waktu
	// Menyisipkan produk baru ke dalam MongoDB
	result, err := r.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return "", err
	}
	// Mendapatkan ID produk yang disimpan
	objectID := result.InsertedID.(primitive.ObjectID)
	// Menghitung durasi waktu yang dihabiskan untuk query
	log.Printf("CreateProduct duration: %v", time.Since(start)) // Mencatat waktu yang diperlukan untuk menyimpan data
	return objectID.Hex(), nil
}

// Mengupdate produk yang sudah ada
func (r *MongoProductRepository) UpdateProduct(product *domain.Product) error {
	start := time.Now() // Mulai pengukuran waktu
	objID, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return err
	}
	// Mengecek apakah produk dengan ID tersebut ada di MongoDB
	if err := r.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Err(); err != nil {
		return err
	}
	// Filter untuk menemukan produk yang akan di-update
	filter := bson.M{"_id": objID}
	// Data yang akan di-update
	update := bson.M{
		"$set": bson.M{
			"name":  product.Name,
			"price": product.Price,
			"stock": product.Stock,
		},
	}
	// Melakukan update pada produk
	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to update product in MongoDB: %v", err)
		return err
	}
	// Menghitung durasi waktu yang dihabiskan untuk update query
	log.Printf("UpdateProduct duration: %v", time.Since(start)) // Mencatat waktu yang dihabiskan untuk proses update
	return nil
}

// Menghapus produk berdasarkan ID
func (r *MongoProductRepository) DeleteProduct(id string) error {
	start := time.Now() // Mulai pengukuran waktu
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// Menghapus produk dari MongoDB berdasarkan ID
	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	// Menghitung durasi waktu yang dihabiskan untuk query penghapusan
	log.Printf("DeleteProduct duration: %v", time.Since(start)) // Mencatat durasi untuk operasi penghapusan
	return nil
}

// Mendapatkan daftar produk
func (r *MongoProductRepository) ListProducts() ([]*domain.Product, error) {
	start := time.Now() // Mulai pengukuran waktu
	// Mengambil semua produk dari MongoDB
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	products := make([]*domain.Product, 0)
	// Iterasi hasil query untuk memasukkan produk ke dalam slice
	for cursor.Next(context.Background()) {
		var product domain.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	// Menghitung durasi waktu yang dihabiskan untuk mengambil semua produk
	log.Printf("ListProducts duration: %v", time.Since(start)) // Mencatat waktu yang dibutuhkan untuk mengambil daftar produk
	return products, nil
}
