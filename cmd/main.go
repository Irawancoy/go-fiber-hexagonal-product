package main

import (
	"context"
	"go-fiber-hexagonal-product/internal/adapters/repositories"
	"go-fiber-hexagonal-product/internal/app"
	"go-fiber-hexagonal-product/pkg/config"
	"go-fiber-hexagonal-product/pkg/database"
	"log"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Inisialisasi repository MongoDB
    mongoClient, err := database.NewMongoDBConnection(cfg.MongoURI)
    if err != nil {
        log.Fatalf("Gagal terhubung ke MongoDB: %v", err)
    }
    defer mongoClient.Disconnect(context.Background())
    
    // Dapatkan koleksi produk dari MongoDB
    mongoCollection := mongoClient.Database(cfg.MongoDatabaseName).Collection("products")
    
    // Buat repository produk MongoDB baru
    mongoRepo := repositories.NewMongoProductRepository(mongoCollection)

    // Inisialisasi repository MySQL
    mysqlDB, err := database.NewMySQLConnection(cfg.MySQLDSN)
    if err != nil {
        log.Fatalf("Gagal terhubung ke MySQL: %v", err)
    }
    defer mysqlDB.Close()
    
    // Buat repository produk MySQL baru
    mysqlRepo := repositories.NewMySQLProductRepository(mysqlDB)

    // Inisialisasi aplikasi dengan kedua repository
    application := app.NewApp(cfg, mongoRepo, mysqlRepo)
    
    // Mulai aplikasi
    log.Fatal(application.Start())
}