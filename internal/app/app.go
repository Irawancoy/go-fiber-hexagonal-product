package app

import (
	"go-fiber-hexagonal-product/internal/adapters/handlers"
	"go-fiber-hexagonal-product/internal/core/ports"
	"go-fiber-hexagonal-product/internal/core/services"
	"go-fiber-hexagonal-product/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	config    *config.Config
	fiberApp  *fiber.App
	mongoRepo ports.MongoProductRepository
	mysqlRepo ports.MySQLProductRepository
}

func NewApp(config *config.Config, mongoRepo ports.MongoProductRepository, mysqlRepo ports.MySQLProductRepository) *App {
	return &App{
		config:    config,
		fiberApp:  fiber.New(),
		mongoRepo: mongoRepo,
		mysqlRepo: mysqlRepo,
	}
}

func (a *App) SetupRoutes() {
	productService := services.NewProductService(a.mongoRepo, a.mysqlRepo)
	productHandler := handlers.NewProductHandler(productService)

	api := a.fiberApp.Group("/api")
	api.Use(logger.New())

	products := api.Group("/products")
	products.Get("/", productHandler.ListProducts)
	products.Post("/", productHandler.CreateProduct)
	products.Get("/:id", productHandler.GetProduct)
	products.Put("/:id", productHandler.UpdateProduct)
	products.Delete("/:id", productHandler.DeleteProduct)
}

func (a *App) Start() error {
	a.SetupRoutes()
	return a.fiberApp.Listen(a.config.ServerAddress)
}
