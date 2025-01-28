package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Auth struct {
	AccessTokenDuration time.Duration
	JwtKey              string
}

func LoadAuth() Auth {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	durationStr := os.Getenv("AUTH_ACCESS_TOKEN_DURATION")
	if durationStr == "" {
		log.Fatal("Missing configuration(s)")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Fatal("Missing configuration(s)")
	}

	jwtKey := os.Getenv("AUTH_JWT_KEY")
	if jwtKey == "" {
		log.Fatal("Missing configuration(s)")
	}

	return Auth{
		AccessTokenDuration: duration,
		JwtKey:              jwtKey,
	}
}
