package database

import (
	"fmt"
	"log"
	"os"

	"github.com/lucsbasto/backend-mineiro/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbHost := os.Getenv("DATABASE_HOST")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbPort := os.Getenv("DATABASE_PORT")

	if dbHost == "" || dbName == "" || dbPassword == "" || dbPort == "" || dbUser  == "" {
		log.Fatal("❌ Algumas variáveis de ambiente não estão definidas! Verifique POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_HOST, POSTGRES_PORT")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost,dbPort, dbUser, dbPassword, dbName)
	fmt.Println( dbHost,dbPort, dbUser, dbPassword, dbName, dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Erro ao conectar ao banco de dados:", err)
	}

	DB = db
	fmt.Println("✅ Conectado ao banco de dados com sucesso!")

	// Executar migrações automaticamente
	err = db.AutoMigrate(&models.Product{}, &models.User{}, &models.Sales{})
	if err != nil {
		log.Fatal("❌ Erro ao rodar migrations:", err)
	}
}
