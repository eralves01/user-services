package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig carrega variáveis de ambiente do arquivo .env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env file not found, loading system variables... %v", err)
	} else {
		log.Println("Environment variables loaded from .env file...")
	}

	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Missing required environment variable: %s", v)
		}
	}
}
