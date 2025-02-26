package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lucsbasto/backend-mineiro/configs"
	"github.com/lucsbasto/backend-mineiro/routes"
)

func main() {
	authController, err := configs.InitializeDependencies()

	if err != nil {
		log.Fatal("❌ Erro ao inicializar dependências: ", err)
	}

	r := gin.Default()
	
	routes.SetupRoutes(r, authController)
	port := os.Getenv("HTTP_PORT")
	r.Run(":"+ port)
}
