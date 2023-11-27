package config

import "github.com/kelseyhightower/envconfig"

type Database struct {
	URI string `required:"true"`
}

func LoadDatabase() Database {
	var config Database
	envconfig.MustProcess("DATABASE", &config)
	return config
}
