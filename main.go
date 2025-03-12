package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/configs"
	"github.com/lucsbasto/backend-mineiro/database"
	"github.com/lucsbasto/backend-mineiro/routes"
)

func main() {
	// configs.Init()
	// Conectar ao banco de dados
	database.Connect()

	// Inicialização das dependências dos controllers
	controllers, err := initializeControllers()
	if err != nil {
		log.Fatal("❌ Erro ao inicializar dependências: ", err)
	}

	// Configuração do gin e das rotas
	r := gin.Default()
	routes.SetupRoutes(r, controllers, database.DB)

	// Iniciando o servidor
	startServer(r)
}

func initializeControllers() (routes.Controllers, error) {
	authController, err := configs.InitializeAuthDependencies()
	if err != nil {
		return routes.Controllers{}, err
	}

	productController, err := configs.InitializeProductDependencies()
	if err != nil {
		return routes.Controllers{}, err
	}

	salesController, err := configs.InitializeSalesDependencies()
	if err != nil {
		return routes.Controllers{}, err
	}

	salesProductController, err := configs.InitializeSalesProductDependencies()
	if err != nil {
		return routes.Controllers{}, err
	}

	return routes.Controllers{
		AuthController:         authController,
		ProductController:      productController,
		SalesController:        salesController,
		SalesProductController: salesProductController,
	}, nil
}

func startServer(r *gin.Engine) {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080" // Valor default se não estiver configurado
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Erro ao iniciar o servidor:", err)
	}
}
