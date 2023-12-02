package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Auth struct {
	AccessTokenDuration time.Duration `envconfig:"ACCESS_TOKEN_DURATION" required:"true"`
	JwtKey              string        `envconfig:"JWT_KEY" required:"true"`
}

func LoadAuth() Auth {
	var config Auth
	envconfig.MustProcess("AUTH", &config)
	return config
}
