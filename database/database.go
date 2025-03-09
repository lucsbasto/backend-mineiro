package database

import (
	"fmt"
	"log"
	"os"

	"github.com/lucsbasto/backend-mineiro/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect estabelece a conexão com o banco de dados.
func Connect() {
	dsn, err := getDSN()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("❌ Erro ao conectar ao banco de dados: ", err)
	}

	DB = db
	fmt.Println("✅ Conectado ao banco de dados com sucesso!")

	err = runMigrations(db)
	if err != nil {
		log.Fatal("❌ Erro ao rodar migrations: ", err)
	}
}

// getDSN constrói a string de conexão DSN a partir das variáveis de ambiente.
func getDSN() (string, error) {
	dbHost := os.Getenv("DATABASE_HOST")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbPort := os.Getenv("DATABASE_PORT")

	if dbHost == "" || dbName == "" || dbPassword == "" || dbPort == "" || dbUser == "" {
		return "", fmt.Errorf("❌ Algumas variáveis de ambiente não estão definidas! Verifique DATABASE_USER, DATABASE_PASSWORD, DATABASE_NAME, DATABASE_HOST, DATABASE_PORT")
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName), nil
}

// runMigrations executa as migrações automáticas para as tabelas do banco de dados.
func runMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Product{}, &models.Sales{}, &models.SalesProduct{})
}
