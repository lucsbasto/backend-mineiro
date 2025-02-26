package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
    err := godotenv.Load()

    if err != nil {
        log.Fatal("❌ Erro ao carregar o arquivo .env")
    }
}