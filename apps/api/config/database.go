package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	URI string `required:"true"`
}

func LoadDatabase() Database {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	uri := os.Getenv("DATABASE_URI")
	if uri == "" {
		log.Fatal("Missing configuration(s)")
	}

	return Database{
		URI: uri,
	}
}
