package configs

import (
	"github.com/lucsbasto/backend-mineiro/controllers"
	"github.com/lucsbasto/backend-mineiro/database"
	"github.com/lucsbasto/backend-mineiro/repositories"
	"github.com/lucsbasto/backend-mineiro/services"
)

func InitializeAuthDependencies() (*controllers.AuthController, error) {
	// Init()
	userRepo := repositories.NewUserRepository(database.DB)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	return authController, nil
}

func InitializeProductDependencies() (*controllers.ProductController, error) {
	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	return productController, nil
}

func InitializeSalesDependencies() (*controllers.SalesController, error) {
	salesRepo := repositories.NewSalesRepository(database.DB)
	productRepo := repositories.NewProductRepository(database.DB)
	salesService := services.NewSalesService(salesRepo, productRepo)
	salesController := controllers.NewSalesController(salesService)

	return salesController, nil
}

func InitializeSalesProductDependencies() (*controllers.SalesProductController, error) {
	salesProductRepo := repositories.NewSalesProductRepository(database.DB)
	salesProductService := services.NewSalesProductService(salesProductRepo)
	salesProductController := controllers.NewSalesProductController(salesProductService)

	return salesProductController, nil
}
